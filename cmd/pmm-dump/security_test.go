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

package main

import (
	"strings"
	"testing"

	"github.com/alecthomas/kingpin/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedactFlagValue(t *testing.T) {
	tests := []struct {
		name  string
		flag  string
		value string
		want  string
	}{
		{
			name:  "pmm token",
			flag:  "pmm-token",
			value: "token-secret",
			want:  "***",
		},
		{
			name:  "pmm cookie",
			flag:  "pmm-cookie",
			value: "cookie-secret",
			want:  "***",
		},
		{
			name:  "encryption pass",
			flag:  "pass",
			value: "dump-secret",
			want:  "***",
		},
		{
			name:  "url credentials",
			flag:  "postgres-url",
			value: "postgres://pg_user:pg_pass@pg.example.com:5432/pmm-managed?username=admin&password=secret",
			want:  "postgres://REDACTED@pg.example.com:5432/pmm-managed?password=REDACTED&username=REDACTED",
		},
		{
			name:  "plain flag",
			flag:  "workers",
			value: "4",
			want:  "4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, redactFlagValue(tt.flag, tt.value))
		})
	}
}

func TestValidateNoSecretCLIArgs(t *testing.T) {
	// Only dedicated secret flags are rejected: they have env-var equivalents and
	// must never appear in process arguments.
	rejected := [][]string{
		{"--pmm-token", "token-secret", "export"},
		{"--pmm-cookie=cookie-secret", "export"},
		{"--pmm-pass", "pmm-secret", "export"},
		{"--pass=dump-secret", "export"},
	}

	for _, args := range rejected {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			require.Error(t, validateNoSecretCLIArgs(args))
		})
	}

	// Credentials embedded in connection-URL flags are allowed: that is the
	// standard, documented pmm-dump invocation (and the form used by the Makefile).
	// They are redacted in logs by redactFlagValue, covered in TestRedactFlagValue.
	accepted := [][]string{
		nil,
		{"--pmm-token=", "export"},
		{"--pmm-url", "http://localhost:8080", "export"},
		{"--pmm-url", "http://admin:secret@localhost:8080", "export"},
		{"--click-house-url=clickhouse://ch.example.com:9000/pmm"},
		{"--click-house-url=clickhouse://ch_user:ch_pass@ch.example.com:9000/pmm"},
		{"--victoria-metrics-url=http://vm.example.com?token=vm-secret"},
		{"--victoria-metrics-url", "http://vm.example.com/prometheus"},
		{"--postgres-url", "postgres://pg.example.com:5432/pmm-managed"},
		{"--postgres-url=postgres://pg_user:pg_pass@pg.example.com:5432/pmm-managed"},
	}

	for _, args := range accepted {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			require.NoError(t, validateNoSecretCLIArgs(args))
		})
	}
}

func TestRedactedArguments(t *testing.T) {
	app := kingpin.New("pmm-dump-test", "")
	app.Flag("pmm-token", "").String()
	app.Flag("pmm-cookie", "").String()
	app.Flag("pass", "").String()
	app.Flag("postgres-url", "").String()
	app.Command("export", "")

	context, err := app.ParseContext([]string{
		"--pmm-token", "token-secret",
		"--pmm-cookie", "cookie-secret",
		"--pass", "dump-secret",
		"--postgres-url", "postgres://pg_user:pg_pass@pg.example.com:5432/pmm-managed?username=admin&password=secret",
		"export",
	})
	require.NoError(t, err)

	args := strings.Join(redactedArguments(context), " ")
	assert.Contains(t, args, "--pmm-token=***")
	assert.Contains(t, args, "--pmm-cookie=***")
	assert.Contains(t, args, "--pass=***")
	assert.Contains(t, args, "--postgres-url=")
	assert.NotContains(t, args, "token-secret")
	assert.NotContains(t, args, "cookie-secret")
	assert.NotContains(t, args, "dump-secret")
	assert.NotContains(t, args, "pg_user")
	assert.NotContains(t, args, "pg_pass")
	assert.NotContains(t, args, "admin")
	assert.NotContains(t, args, "secret")
}
