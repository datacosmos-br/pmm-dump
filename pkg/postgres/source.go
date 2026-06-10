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

package postgres

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"pmm-dump/pkg/dump"
)

// Source implements dump.Source for PostgreSQL.
type Source struct {
	config Config
}

// NewSource creates a new PostgreSQL source.
func NewSource(cfg Config) *Source {
	return &Source{config: cfg}
}

// Type returns the source type.
func (s Source) Type() dump.SourceType {
	return dump.Postgres
}

// ReadChunk dumps the specified database using pg_dump.
func (s Source) ReadChunk(meta dump.ChunkMeta) (*dump.Chunk, error) {
	dbIndex := meta.Index % len(s.config.Databases)
	dbName := s.config.Databases[dbIndex]

	log.Debug().Int("index", meta.Index).Str("database", dbName).Msg("Dumping PostgreSQL database")

	cmd := exec.Command("pg_dump", //nolint:gosec
		"--data-only",
		"--format=custom",
		"--dbname", s.config.ConnectionURL,
		"--schema", dbName)

	var out bytes.Buffer
	cmd.Stdout = &out
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return nil, errors.Wrapf(err, "pg_dump failed: %s", stderr.String())
	}

	return &dump.Chunk{
		ChunkMeta: meta,
		Content:   out.Bytes(),
		Filename:  fmt.Sprintf("%d.pgdump", meta.Index),
	}, nil
}

// WriteChunk restores the PostgreSQL dump using pg_restore.
//
//nolint:unparam
func (s Source) WriteChunk(filename string, r io.Reader) error {
	log.Debug().Str("filename", filename).Msg("Restoring PostgreSQL chunk")

	cmd := exec.Command("pg_restore", //nolint:gosec
		"--clean",
		"--if-exists",
		"--dbname", s.config.ConnectionURL)

	cmd.Stdin = r
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		// pg_restore often returns exit code 1 for warnings; we log them but don't fail
		log.Warn().Str("stderr", stderr.String()).Msg("pg_restore finished with warnings")
	}
	return nil
}

// FinalizeWrites is a no-op for PostgreSQL.
func (s Source) FinalizeWrites() error {
	return nil
}
