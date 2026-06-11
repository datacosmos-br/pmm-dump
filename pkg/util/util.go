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
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/hashicorp/go-version"
)

type PMMConfig struct {
	PMMURL             string
	ClickHouseURL      string
	VictoriaMetricsURL string
}

func GetPMMConfig(pmmLink, vmLink, chLink string, ver *version.Version) (PMMConfig, error) {
	pmmURL, err := url.Parse(pmmLink)
	if err != nil {
		return PMMConfig{}, fmt.Errorf("failed to parse pmm-url: %w", err)
	}
	conf := PMMConfig{
		PMMURL:             pmmLink,
		ClickHouseURL:      chLink,
		VictoriaMetricsURL: vmLink,
	}

	if conf.ClickHouseURL == "" {
		conf.ClickHouseURL = clickHouseURLFromEnvComponents()
	}

	if conf.ClickHouseURL == "" {
		conf.ClickHouseURL = composeClickHouseURL(*pmmURL, ver)
	}
	if conf.VictoriaMetricsURL == "" {
		conf.VictoriaMetricsURL = composeVictoriaMetricsURL(*pmmURL)
	}
	return conf, nil
}

func composeVictoriaMetricsURL(u url.URL) string {
	u.Path = "/prometheus"
	u.RawQuery = ""
	return u.String()
}

func composeClickHouseURL(u url.URL, ver *version.Version) string {
	u.Scheme = "clickhouse"
	i := strings.LastIndex(u.Host, ":")
	if i != -1 {
		u.Host = u.Host[:i]
	}

	u.User = url.UserPassword("default", "clickhouse")
	if CheckVer(ver, "<= 3.1.0") {
		u.User = nil
	}

	u.Host += ":9000"
	u.Path = "pmm"
	return u.String()
}

func CheckVer(ver *version.Version, constrain string) bool {
	if ver == nil {
		return false
	}
	constraints, err := version.NewConstraint(constrain)
	if err != nil {
		panic(fmt.Sprintf("cannot create constraint: %v", err))
	}
	resConst := ver
	return constraints.Check(resConst)
}

func clickHouseURLFromEnvComponents() string {
	addr := os.Getenv("PMM_CLICKHOUSE_ADDR")
	if addr == "" {
		host := os.Getenv("PMM_CLICKHOUSE_HOST")
		if host == "" {
			return ""
		}
		port := os.Getenv("PMM_CLICKHOUSE_PORT")
		if port != "" {
			addr = net.JoinHostPort(host, port)
		} else {
			addr = host
		}
	}
	user := os.Getenv("PMM_CLICKHOUSE_USER")
	pass := os.Getenv("PMM_CLICKHOUSE_PASSWORD")
	db := os.Getenv("PMM_CLICKHOUSE_DATABASE")

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

// RedactURL removes credentials from a URL for safe logging.
func RedactURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return "***"
	}
	if u.User != nil {
		u.User = url.User("REDACTED")
	}
	query := u.Query()
	for key := range query {
		lowerKey := strings.ToLower(key)
		if lowerKey == "user" || lowerKey == "username" || strings.Contains(lowerKey, "pass") || strings.Contains(lowerKey, "token") || strings.Contains(lowerKey, "secret") {
			query.Set(key, "REDACTED")
		}
	}
	u.RawQuery = query.Encode()
	if u.String() == "" {
		return "***"
	}
	return u.String()
}
