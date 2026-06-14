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

package victoriametrics

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/valyala/fasthttp"

	"pmm-dump/pkg/dump"
	"pmm-dump/pkg/grafana/client"
)

func TestWriteChunk(t *testing.T) {
	tests := []struct {
		name         string
		metricsSize  int
		contentLimit int
		nativeData   bool
		shouldErr    bool
	}{
		{
			name:         "native data",
			metricsSize:  20,
			contentLimit: 20,
			nativeData:   true,
			shouldErr:    true,
		},
		{
			name:         "0 content limit",
			metricsSize:  20,
			contentLimit: 0,
		},
		{
			name:         "with content limit",
			metricsSize:  20,
			contentLimit: 130,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpC := &fasthttp.Client{
				MaxConnsPerHost:           2,
				MaxIdleConnDuration:       time.Minute,
				MaxIdemponentCallAttempts: 5,
				ReadTimeout:               time.Minute,
				WriteTimeout:              time.Minute,
				MaxConnWaitTimeout:        time.Second * 30,
				TLSConfig: &tls.Config{
					InsecureSkipVerify: true, //nolint:gosec
				},
			}
			var recievedMetrics []Metric
			server := httptest.NewServer(http.HandlerFunc(
				func(rw http.ResponseWriter, req *http.Request) {
					defer req.Body.Close() //nolint:errcheck
					if req.ContentLength > int64(tt.contentLimit) && tt.contentLimit != 0 {
						rw.WriteHeader(http.StatusRequestEntityTooLarge)
						return
					}
					compressedContent, err := io.ReadAll(req.Body)
					if err != nil {
						t.Error(err)
						rw.WriteHeader(http.StatusBadRequest)
						return
					}
					metrics, err := decompressChunk(compressedContent)
					if err != nil {
						t.Error(err)
						rw.WriteHeader(http.StatusBadRequest)
						return
					}
					recievedMetrics = append(recievedMetrics, metrics...)
				},
			))
			defer server.Close()

			grafanaC, err := client.NewClient(httpC, client.AuthParams{
				User:     "admin",
				Password: "admin",
			})
			if err != nil {
				t.Fatal(err)
			}
			s := NewSource(grafanaC, &Config{
				ConnectionURL: server.URL,
				NativeData:    tt.nativeData,
				ContentLimit:  tt.contentLimit,
			})

			data, err := generateFakeChunk(tt.metricsSize)
			if err != nil {
				t.Fatal(err)
			}
			err = s.WriteChunk("", bytes.NewBuffer(data))
			if err != nil && !tt.shouldErr {
				t.Fatal(err)
			}
			if err == nil && tt.shouldErr {
				t.Fatal("should be error")
			}
		})
	}
}

func TestReadChunksNormalizesUncompressedVMResponse(t *testing.T) {
	// Root-cause regression: VictoriaMetrics honours Accept-Encoding: gzip only
	// above an internal size threshold, so tiny boundary chunks come back as
	// raw JSON. The dump format requires every VM chunk to be gzip — the import
	// (sendChunk) and split (decompressChunk) paths assume it — so ReadChunks
	// must normalise a raw response to gzip before storing it. Without this,
	// import fails with "cannot decode vmimport data: gzip: invalid header".
	const rawLine = `{"metric":{"__name__":"node_test","instance":"x"},"values":[0],"timestamps":[1781461557000]}` + "\n"

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, _ *http.Request) {
			// Respond uncompressed regardless of Accept-Encoding, mimicking VM
			// skipping gzip for a small payload.
			_, _ = io.WriteString(rw, rawLine)
		},
	))
	defer server.Close()

	httpC := &fasthttp.Client{
		MaxConnsPerHost:           2,
		MaxIdleConnDuration:       time.Minute,
		MaxIdemponentCallAttempts: 5,
		ReadTimeout:               time.Minute,
		WriteTimeout:              time.Minute,
		MaxConnWaitTimeout:        time.Second * 30,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec
		},
	}
	grafanaC, err := client.NewClient(httpC, client.AuthParams{
		User:     "admin",
		Password: "admin",
	})
	if err != nil {
		t.Fatal(err)
	}
	s := NewSource(grafanaC, &Config{ConnectionURL: server.URL})

	start := time.Unix(1781461557, 0)
	end := time.Unix(1781461587, 0)
	chunks, err := s.ReadChunks(dump.ChunkMeta{Source: dump.VictoriaMetrics, Start: &start, End: &end})
	if err != nil {
		t.Fatalf("ReadChunks: %v", err)
	}
	if len(chunks) != 1 {
		t.Fatalf("want 1 chunk, got %d", len(chunks))
	}

	// Invariant: stored chunk content must be valid gzip.
	if _, err := gzip.NewReader(bytes.NewReader(chunks[0].Content)); err != nil {
		t.Fatalf("chunk content is not gzip (dump invariant violated): %v", err)
	}

	// And it must round-trip back to the original metric.
	metrics, err := decompressChunk(chunks[0].Content)
	if err != nil {
		t.Fatalf("decompressChunk: %v", err)
	}
	if len(metrics) != 1 || metrics[0].Metric["__name__"] != "node_test" {
		t.Fatalf("unexpected metrics round-trip: %+v", metrics)
	}
}

func generateFakeChunk(size int) ([]byte, error) {
	metricsData, err := json.Marshal(Metric{
		Metric: map[string]string{
			"__name__": "test",
			"job":      "test",
			"instance": "test",
			"test":     "test",
		},
		Values:     []float64{100000000000000},
		Timestamps: []int64{time.Now().Unix()},
	})
	if err != nil {
		return nil, fmt.Errorf("marshal metrics: %w", err)
	}
	var data []byte
	for range size {
		data = append(data, metricsData...)
	}
	return compressData(data)
}

func compressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(data); err != nil {
		return nil, fmt.Errorf("write gzip: %w", err)
	}
	if err := gw.Close(); err != nil {
		return nil, fmt.Errorf("close gzip: %w", err)
	}
	return buf.Bytes(), nil
}
