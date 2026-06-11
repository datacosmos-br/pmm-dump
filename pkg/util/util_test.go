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

package util

import (
	"os"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClickHouseURLFromEnv(t *testing.T) {
	t.Run("complete URL", func(t *testing.T) {
		t.Setenv("PMM_CLICKHOUSE_URL", "clickhouse://user:pass@host:9000/db")
		assert.Equal(t, "clickhouse://user:pass@host:9000/db", GetClickHouseURLFromEnv())
	})

	t.Run("build from components", func(t *testing.T) {
		t.Setenv("PMM_CLICKHOUSE_URL", "")
		t.Setenv("PMM_CLICKHOUSE_ADDR", "ch.example.com:9000")
		t.Setenv("PMM_CLICKHOUSE_USER", "chuser")
		t.Setenv("PMM_CLICKHOUSE_PASSWORD", "chpass")
		t.Setenv("PMM_CLICKHOUSE_DATABASE", "chdb")
		assert.Equal(t, "clickhouse://chuser:chpass@ch.example.com:9000/chdb", GetClickHouseURLFromEnv())
	})

	t.Run("default", func(t *testing.T) {
		_ = os.Unsetenv("PMM_CLICKHOUSE_URL")
		_ = os.Unsetenv("PMM_CLICKHOUSE_ADDR")
		_ = os.Unsetenv("PMM_CLICKHOUSE_USER")
		_ = os.Unsetenv("PMM_CLICKHOUSE_PASSWORD")
		_ = os.Unsetenv("PMM_CLICKHOUSE_DATABASE")
		assert.Equal(t, "clickhouse://default:clickhouse@127.0.0.1:9000/pmm", GetClickHouseURLFromEnv())
	})
}

func TestGetVMURLFromEnv(t *testing.T) {
	t.Run("from env", func(t *testing.T) {
		t.Setenv("PMM_VM_URL", "http://vm:8428")
		assert.Equal(t, "http://vm:8428", GetVMURLFromEnv())
	})

	t.Run("default", func(t *testing.T) {
		_ = os.Unsetenv("PMM_VM_URL")
		assert.Equal(t, "http://127.0.0.1:9090/prometheus", GetVMURLFromEnv())
	})
}

func TestGetPMMConfig(t *testing.T) {
	t.Run("all explicit", func(t *testing.T) {
		ver, _ := version.NewVersion("3.8.1")
		conf, err := GetPMMConfig("http://pmm:8080", "http://vm:8428", "clickhouse://ch:9000", "postgres://pg:5432/pmm", ver)
		require.NoError(t, err)
		assert.Equal(t, "http://pmm:8080", conf.PMMURL)
		assert.Equal(t, "http://vm:8428", conf.VictoriaMetricsURL)
		assert.Equal(t, "clickhouse://ch:9000", conf.ClickHouseURL)
		assert.Equal(t, "postgres://pg:5432/pmm", conf.PostgresURL)
	})

	t.Run("fallback to env", func(t *testing.T) {
		t.Setenv("PMM_VM_URL", "http://vm-env:8428")
		t.Setenv("PMM_CLICKHOUSE_URL", "clickhouse://ch-env:9000")
		t.Setenv("PMM_POSTGRES_URL", "postgres://pg-env:5432/pmm")
		ver, _ := version.NewVersion("3.8.1")
		conf, err := GetPMMConfig("http://pmm:8080", "", "", "", ver)
		require.NoError(t, err)
		assert.Equal(t, "http://vm-env:8428", conf.VictoriaMetricsURL)
		assert.Equal(t, "clickhouse://ch-env:9000", conf.ClickHouseURL)
		assert.Equal(t, "postgres://pg-env:5432/pmm", conf.PostgresURL)
	})
}

func TestGetPostgresURLFromEnv(t *testing.T) {
	t.Run("empty without env", func(t *testing.T) {
		assert.Empty(t, GetPostgresURLFromEnv())
	})

	t.Run("explicit url", func(t *testing.T) {
		t.Setenv("PMM_POSTGRES_URL", "postgres://pg.example.com:5432/pmm-managed")
		assert.Equal(t, "postgres://pg.example.com:5432/pmm-managed", GetPostgresURLFromEnv())
	})

	t.Run("component env", func(t *testing.T) {
		t.Setenv("PMM_POSTGRES_ADDR", "pg.example.com:5432")
		t.Setenv("PMM_POSTGRES_USERNAME", "pmm")
		t.Setenv("PMM_POSTGRES_PASSWORD", "x")
		t.Setenv("PMM_POSTGRES_DBNAME", "pmm-managed")
		t.Setenv("PMM_POSTGRES_SSLMODE", "require")

		assert.Equal(t, "postgres://pmm:x@pg.example.com:5432/pmm-managed?sslmode=require", GetPostgresURLFromEnv())
	})
}

func TestRedactURL(t *testing.T) {
	t.Run("redact credentials", func(t *testing.T) {
		redacted := RedactURL("clickhouse://user:pass@host:9000/db")
		assert.Contains(t, redacted, "REDACTED")
		assert.NotContains(t, redacted, "pass")
	})

	t.Run("no credentials", func(t *testing.T) {
		redacted := RedactURL("http://host:8080")
		assert.Equal(t, "http://host:8080", redacted)
	})
}
