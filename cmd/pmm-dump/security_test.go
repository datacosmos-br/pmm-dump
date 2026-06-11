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
			flag:  "click-house-url",
			value: "clickhouse://ch_user:ch_pass@ch.example.com:9000/pmm?username=admin&password=secret",
			want:  "clickhouse://REDACTED@ch.example.com:9000/pmm?password=REDACTED&username=REDACTED",
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
	rejected := [][]string{
		{"--pmm-token", "token-secret", "export"},
		{"--pmm-cookie=cookie-secret", "export"},
		{"--pmm-pass", "pmm-secret", "export"},
		{"--pass=dump-secret", "export"},
		{"--pmm-url", "http://admin:secret@localhost:8080", "export"},
		{"--click-house-url=clickhouse://ch_user:ch_pass@ch.example.com:9000/pmm"},
		{"--victoria-metrics-url=http://vm.example.com?token=vm-secret"},
	}

	for _, args := range rejected {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			require.Error(t, validateNoSecretCLIArgs(args))
		})
	}

	accepted := [][]string{
		nil,
		{"--pmm-token=", "export"},
		{"--pmm-url", "http://localhost:8080", "export"},
		{"--click-house-url=clickhouse://ch.example.com:9000/pmm"},
		{"--victoria-metrics-url", "http://vm.example.com/prometheus"},
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
	app.Flag("click-house-url", "").String()
	app.Command("export", "")

	context, err := app.ParseContext([]string{
		"--pmm-token", "token-secret",
		"--pmm-cookie", "cookie-secret",
		"--pass", "dump-secret",
		"--click-house-url", "clickhouse://ch_user:ch_pass@ch.example.com:9000/pmm?username=admin&password=secret",
		"export",
	})
	require.NoError(t, err)

	args := strings.Join(redactedArguments(context), " ")
	assert.Contains(t, args, "--pmm-token=***")
	assert.Contains(t, args, "--pmm-cookie=***")
	assert.Contains(t, args, "--pass=***")
	assert.Contains(t, args, "--click-house-url=")
	assert.NotContains(t, args, "token-secret")
	assert.NotContains(t, args, "cookie-secret")
	assert.NotContains(t, args, "dump-secret")
	assert.NotContains(t, args, "ch_user")
	assert.NotContains(t, args, "ch_pass")
	assert.NotContains(t, args, "admin")
	assert.NotContains(t, args, "secret")
}
