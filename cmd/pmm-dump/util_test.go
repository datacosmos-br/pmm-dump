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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"pmm-dump/pkg/dump"
	"pmm-dump/pkg/util"
)

func TestRedactURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "with credentials",
			input:    "http://user:password@localhost:8080/path",
			expected: "http://REDACTED:REDACTED@localhost:8080/path",
		},
		{
			name:     "without credentials",
			input:    "http://localhost:8080/path",
			expected: "http://localhost:8080/path",
		},
		{
			name:     "invalid URL without credentials",
			input:    "://invalid-url",
			expected: "://invalid-url",
		},
		{
			name:     "postgres URL",
			input:    "postgres://admin:secret@db.example.com:5432/pmm-managed?sslmode=require",
			expected: "postgres://REDACTED:REDACTED@db.example.com:5432/pmm-managed?sslmode=require",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, util.RedactURL(tt.input))
		})
	}
}

func TestPreparePostgresSource(t *testing.T) {
	t.Run("enabled with URL", func(t *testing.T) {
		src, ok := preparePostgresSource(true, "postgres://localhost:5432")
		require.True(t, ok)
		assert.Equal(t, dump.Postgres, src.Type())
	})

	t.Run("enabled without URL", func(t *testing.T) {
		_, ok := preparePostgresSource(true, "")
		assert.False(t, ok)
	})

	t.Run("disabled", func(t *testing.T) {
		_, ok := preparePostgresSource(false, "postgres://localhost:5432")
		assert.False(t, ok)
	})
}
