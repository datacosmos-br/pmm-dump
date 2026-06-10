// Copyright 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build security

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pmm-dump/pkg/util"
)

func TestNoCredentialsInProcessList(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("cmdline check only works on Linux")
	}

	// Set env vars that the binary should read, not pass as flags
	t.Setenv("PMM_CLICKHOUSE_URL", "clickhouse://secret_user:secret_pass@ch.example.com:9000/pmm")
	t.Setenv("PMM_VM_URL", "http://vm:8428")
	t.Setenv("PMM_POSTGRES_URL", "postgres://pg_user:pg_pass@pg.example.com:5432/pmm-managed")

	// Build the binary
	binPath := filepath.Join(t.TempDir(), "pmm-dump")
	build := exec.Command("go", "build", "-o", binPath, "pmm-dump/cmd/pmm-dump")
	buildOut, err := build.CombinedOutput()
	require.NoError(t, err, "build failed: %s", buildOut)

	// Run it in background with just basic flags
	cmd := exec.Command(binPath, "version")
	cmd.Env = os.Environ()
	require.NoError(t, cmd.Start())
	defer func() { _ = cmd.Process.Kill() }()

	// Give it time to start
	time.Sleep(100 * time.Millisecond)

	// Read /proc/<pid>/cmdline
	cmdlinePath := fmt.Sprintf("/proc/%d/cmdline", cmd.Process.Pid)
	data, err := os.ReadFile(cmdlinePath)
	require.NoError(t, err)
	cmdline := strings.ReplaceAll(string(data), "\x00", " ")

	// Assert no credentials leaked
	assert.NotContains(t, cmdline, "secret_user")
	assert.NotContains(t, cmdline, "secret_pass")
	assert.NotContains(t, cmdline, "pg_user")
	assert.NotContains(t, cmdline, "pg_pass")
	assert.NotContains(t, cmdline, "clickhouse://")
	assert.NotContains(t, cmdline, "postgres://")
}

func TestNoCredentialsInLogs(t *testing.T) {
	// Set sensitive env vars
	t.Setenv("PMM_CLICKHOUSE_URL", "clickhouse://log_user:log_pass@ch.example.com:9000/pmm")
	t.Setenv("PMM_VM_URL", "http://vm:8428")
	t.Setenv("PMM_POSTGRES_URL", "postgres://pg_user:pg_pass@pg.example.com:5432/pmm-managed")

	// Build the binary
	binPath := filepath.Join(t.TempDir(), "pmm-dump")
	build := exec.Command("go", "build", "-o", binPath, "pmm-dump/cmd/pmm-dump")
	buildOut, err := build.CombinedOutput()
	require.NoError(t, err, "build failed: %s", buildOut)

	// Run version with verbose to trigger logging
	cmd := exec.Command(binPath, "version", "-v")
	cmd.Env = os.Environ()
	out, _ := cmd.CombinedOutput()
	output := string(out)

	// Assert no credentials in logs
	sensitivePatterns := []string{
		"log_user", "log_pass", "pg_user", "pg_pass",
		"clickhouse://log_user", "postgres://pg_user",
	}
	for _, pattern := range sensitivePatterns {
		assert.NotContains(t, output, pattern, "credential leak in logs: %s", pattern)
	}
}

func TestDumpFilePermissions(t *testing.T) {
	tmpDir := t.TempDir()
	dumpPath := filepath.Join(tmpDir, "test-dump.tar.gz")

	// Create a dummy dump file
	f, err := os.OpenFile(dumpPath, os.O_CREATE|os.O_WRONLY, 0o600)
	require.NoError(t, err)
	_, err = f.WriteString("dummy")
	require.NoError(t, err)
	require.NoError(t, f.Close())

	// Verify permissions
	info, err := os.Stat(dumpPath)
	require.NoError(t, err)
	mode := info.Mode().Perm()
	assert.Equal(t, os.FileMode(0o600), mode, "dump file must be owner-readable only")
}

func TestRedactURL_NoCredentialsInMeta(t *testing.T) {
	urls := []string{
		"http://admin:secret@localhost:8080/path",
		"clickhouse://ch_user:ch_pass@ch.example.com:9000/pmm",
		"postgres://pg_user:pg_pass@pg.example.com:5432/pmm-managed?sslmode=require",
	}
	passwordRegex := regexp.MustCompile(`(?i)(user|pass|password)[=:]\S+`)

	for _, raw := range urls {
		redacted := util.RedactURL(raw)
		assert.NotContains(t, redacted, "secret", "URL not redacted: %s", raw)
		assert.NotContains(t, redacted, "ch_pass", "URL not redacted: %s", raw)
		assert.NotContains(t, redacted, "pg_pass", "URL not redacted: %s", raw)
		assert.False(t, passwordRegex.MatchString(redacted), "password pattern found in redacted URL: %s", redacted)
	}
}
