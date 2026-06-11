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
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClickHouseURLFromEnvComponents(t *testing.T) {
	t.Run("build from components", func(t *testing.T) {
		clearClickHouseEnv(t)
		t.Setenv("PMM_CLICKHOUSE_ADDR", "ch.example.com:9000")
		t.Setenv("PMM_CLICKHOUSE_USER", "chuser")
		t.Setenv("PMM_CLICKHOUSE_PASSWORD", "chpass")
		t.Setenv("PMM_CLICKHOUSE_DATABASE", "chdb")
		assert.Equal(t, "clickhouse://chuser:chpass@ch.example.com:9000/chdb", clickHouseURLFromEnvComponents())
	})

	t.Run("host and port components", func(t *testing.T) {
		clearClickHouseEnv(t)
		t.Setenv("PMM_CLICKHOUSE_HOST", "ch.example.com")
		t.Setenv("PMM_CLICKHOUSE_PORT", "9000")
		t.Setenv("PMM_CLICKHOUSE_USER", "chuser")
		t.Setenv("PMM_CLICKHOUSE_PASSWORD", "chpass")
		t.Setenv("PMM_CLICKHOUSE_DATABASE", "chdb")
		assert.Equal(t, "clickhouse://chuser:chpass@ch.example.com:9000/chdb", clickHouseURLFromEnvComponents())
	})

	t.Run("empty when env is absent", func(t *testing.T) {
		clearClickHouseEnv(t)
		assert.Empty(t, clickHouseURLFromEnvComponents())
	})
}

func TestGetPMMConfig(t *testing.T) {
	t.Run("all explicit", func(t *testing.T) {
		ver, err := version.NewVersion("3.8.1")
		require.NoError(t, err)
		conf, err := GetPMMConfig("http://pmm:8080", "http://vm:8428", "clickhouse://ch:9000", "postgres://pg:5432/pmm", ver)
		require.NoError(t, err)
		assert.Equal(t, "http://pmm:8080", conf.PMMURL)
		assert.Equal(t, "http://vm:8428", conf.VictoriaMetricsURL)
		assert.Equal(t, "clickhouse://ch:9000", conf.ClickHouseURL)
		assert.Equal(t, "postgres://pg:5432/pmm", conf.PostgresURL)
	})

	t.Run("values supplied by kingpin env binding", func(t *testing.T) {
		ver, err := version.NewVersion("3.8.1")
		require.NoError(t, err)
		conf, err := GetPMMConfig("http://pmm:8080", "http://vm-env:8428", "clickhouse://ch-env:9000", "postgres://pg-env:5432/pmm", ver)
		require.NoError(t, err)
		assert.Equal(t, "http://vm-env:8428", conf.VictoriaMetricsURL)
		assert.Equal(t, "clickhouse://ch-env:9000", conf.ClickHouseURL)
		assert.Equal(t, "postgres://pg-env:5432/pmm", conf.PostgresURL)
	})

	t.Run("component env fallback", func(t *testing.T) {
		clearClickHouseEnv(t)
		clearPostgresEnv(t)
		t.Setenv("PMM_CLICKHOUSE_ADDR", "ch-env.example.com:9000")
		t.Setenv("PMM_CLICKHOUSE_USER", "chuser")
		t.Setenv("PMM_CLICKHOUSE_PASSWORD", "chpass")
		t.Setenv("PMM_CLICKHOUSE_DATABASE", "chdb")
		t.Setenv("PMM_POSTGRES_ADDR", "pg-env.example.com:5432")
		t.Setenv("PMM_POSTGRES_USERNAME", "pguser")
		t.Setenv("PMM_POSTGRES_PASSWORD", "pgpass")
		t.Setenv("PMM_POSTGRES_DBNAME", "pgdb")
		t.Setenv("PMM_POSTGRES_SSLMODE", "require")
		ver, err := version.NewVersion("3.8.1")
		require.NoError(t, err)
		conf, err := GetPMMConfig("http://pmm:8080", "", "", "", ver)
		require.NoError(t, err)
		assert.Equal(t, "http://pmm:8080/prometheus", conf.VictoriaMetricsURL)
		assert.Equal(t, "clickhouse://chuser:chpass@ch-env.example.com:9000/chdb", conf.ClickHouseURL)
		assert.Equal(t, "postgres://pguser:pgpass@pg-env.example.com:5432/pgdb?sslmode=require", conf.PostgresURL)
	})

	t.Run("fallback to pmm url when env absent", func(t *testing.T) {
		clearClickHouseEnv(t)
		clearPostgresEnv(t)
		ver, err := version.NewVersion("3.8.1")
		require.NoError(t, err)
		conf, err := GetPMMConfig("http://pmm:8080", "", "", "", ver)
		require.NoError(t, err)
		assert.Equal(t, "http://pmm:8080/prometheus", conf.VictoriaMetricsURL)
		assert.Equal(t, "clickhouse://default:clickhouse@pmm:9000/pmm", conf.ClickHouseURL)
		assert.Empty(t, conf.PostgresURL)
	})
}

func TestPostgresURLFromEnvComponents(t *testing.T) {
	t.Run("empty without env", func(t *testing.T) {
		clearPostgresEnv(t)
		assert.Empty(t, postgresURLFromEnvComponents())
	})

	t.Run("component env", func(t *testing.T) {
		clearPostgresEnv(t)
		t.Setenv("PMM_POSTGRES_ADDR", "pg.example.com:5432")
		t.Setenv("PMM_POSTGRES_USERNAME", "pmm")
		t.Setenv("PMM_POSTGRES_PASSWORD", "x")
		t.Setenv("PMM_POSTGRES_DBNAME", "pmm-managed")
		t.Setenv("PMM_POSTGRES_SSLMODE", "require")

		assert.Equal(t, "postgres://pmm:x@pg.example.com:5432/pmm-managed?sslmode=require", postgresURLFromEnvComponents())
	})
}

func clearClickHouseEnv(t *testing.T) {
	t.Helper()
	for _, key := range []string{
		"PMM_CLICKHOUSE_ADDR",
		"PMM_CLICKHOUSE_HOST",
		"PMM_CLICKHOUSE_PORT",
		"PMM_CLICKHOUSE_USER",
		"PMM_CLICKHOUSE_PASSWORD",
		"PMM_CLICKHOUSE_DATABASE",
	} {
		t.Setenv(key, "")
	}
}

func clearPostgresEnv(t *testing.T) {
	t.Helper()
	for _, key := range []string{
		"PMM_POSTGRES_ADDR",
		"PMM_POSTGRES_USERNAME",
		"PMM_POSTGRES_PASSWORD",
		"PMM_POSTGRES_DBNAME",
		"PMM_POSTGRES_SSLMODE",
	} {
		t.Setenv(key, "")
	}
}

func TestRedactURL(t *testing.T) {
	t.Run("redact credentials", func(t *testing.T) {
		redacted := RedactURL("clickhouse://user:pass@host:9000/db?password=secret&token=abc&username=admin")
		assert.Contains(t, redacted, "REDACTED")
		assert.NotContains(t, redacted, "user:pass")
		assert.NotContains(t, redacted, "secret")
		assert.NotContains(t, redacted, "abc")
		assert.NotContains(t, redacted, "admin")
	})

	t.Run("no credentials", func(t *testing.T) {
		redacted := RedactURL("http://host:8080")
		assert.Equal(t, "http://host:8080", redacted)
	})
}
