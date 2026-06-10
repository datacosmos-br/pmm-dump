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
	"fmt"
	"net/url"
	"os"
	"strings"
)

const defaultPMMDB = "pmm"

type PMMConfig struct {
	PMMURL             string
	ClickHouseURL      string
	VictoriaMetricsURL string
	PostgresURL        string
}

func GetPMMConfig(pmmLink, vmLink, chLink, pgLink string) (PMMConfig, error) {
	pmmURL, err := url.Parse(pmmLink)
	if err != nil {
		return PMMConfig{}, fmt.Errorf("failed to parse pmm-url: %w", err)
	}
	conf := PMMConfig{
		PMMURL:             pmmLink,
		ClickHouseURL:      chLink,
		VictoriaMetricsURL: vmLink,
		PostgresURL:        pgLink,
	}

	if conf.ClickHouseURL == "" {
		conf.ClickHouseURL = GetClickHouseURLFromEnv()
	}
	if conf.VictoriaMetricsURL == "" {
		conf.VictoriaMetricsURL = GetVMURLFromEnv()
	}
	if conf.PostgresURL == "" {
		conf.PostgresURL = GetPostgresURLFromEnv()
	}

	// Fallback to composing from pmm-url if still empty
	if conf.ClickHouseURL == "" {
		conf.ClickHouseURL = composeClickHouseURL(*pmmURL)
	}
	if conf.VictoriaMetricsURL == "" {
		conf.VictoriaMetricsURL = composeVictoriaMetricsURL(*pmmURL)
	}
	return conf, nil
}

// GetClickHouseURLFromEnv builds ClickHouse URL from PMM environment variables.
// Priority: PMM_CLICKHOUSE_URL > PMM_CLICKHOUSE_* individual vars.
func GetClickHouseURLFromEnv() string {
	urlStr := os.Getenv("PMM_CLICKHOUSE_URL")
	if urlStr != "" {
		return urlStr
	}

	addr := os.Getenv("PMM_CLICKHOUSE_ADDR")
	if addr == "" {
		addr = "127.0.0.1:9000"
	}
	user := os.Getenv("PMM_CLICKHOUSE_USER")
	if user == "" {
		user = "default"
	}
	pass := os.Getenv("PMM_CLICKHOUSE_PASSWORD")
	if pass == "" {
		pass = "clickhouse"
	}
	db := os.Getenv("PMM_CLICKHOUSE_DATABASE")
	if db == "" {
		db = defaultPMMDB
	}

	u := url.URL{
		Scheme: "clickhouse",
		Host:   addr,
		Path:   db,
	}
	if user != "" || pass != "" {
		u.User = url.UserPassword(user, pass)
	}
	return u.String()
}

// GetVMURLFromEnv returns VictoriaMetrics URL from PMM_VM_URL env var.
func GetVMURLFromEnv() string {
	urlStr := os.Getenv("PMM_VM_URL")
	if urlStr != "" {
		return urlStr
	}
	return "http://127.0.0.1:9090/prometheus"
}

// GetPostgresURLFromEnv builds PostgreSQL URL from PMM environment variables.
// Priority: PMM_POSTGRES_URL > PMM_POSTGRES_* individual vars.
func GetPostgresURLFromEnv() string {
	urlStr := os.Getenv("PMM_POSTGRES_URL")
	if urlStr != "" {
		return urlStr
	}

	addr := os.Getenv("PMM_POSTGRES_ADDR")
	if addr == "" {
		addr = "127.0.0.1:5432"
	}
	user := os.Getenv("PMM_POSTGRES_USERNAME")
	if user == "" {
		user = defaultPMMDB
	}
	pass := os.Getenv("PMM_POSTGRES_PASSWORD")
	if pass == "" {
		pass = defaultPMMDB
	}
	db := os.Getenv("PMM_POSTGRES_DBNAME")
	if db == "" {
		db = "pmm-managed"
	}
	sslmode := os.Getenv("PMM_POSTGRES_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	u := url.URL{
		Scheme: "postgres",
		Host:   addr,
		Path:   db,
	}
	if user != "" || pass != "" {
		u.User = url.UserPassword(user, pass)
	}
	q := u.Query()
	q.Set("sslmode", sslmode)
	u.RawQuery = q.Encode()
	return u.String()
}

func composeVictoriaMetricsURL(u url.URL) string {
	u.Path = "/prometheus"
	u.RawQuery = ""
	return u.String()
}

func composeClickHouseURL(u url.URL) string {
	u.Scheme = "clickhouse"
	i := strings.LastIndex(u.Host, ":")
	if i != -1 {
		u.Host = u.Host[:i]
	}
	u.User = nil
	u.Host += ":9000"
	u.Path = defaultPMMDB
	return u.String()
}
