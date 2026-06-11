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

func (s Source) Chunks() []dump.ChunkMeta {
	chunks := make([]dump.ChunkMeta, 0, len(s.config.Databases))
	for i := range s.config.Databases {
		chunks = append(chunks, dump.ChunkMeta{
			Source:  dump.Postgres,
			Index:   i,
			RowsLen: 1,
		})
	}
	return chunks
}

func (s Source) ReadChunks(meta dump.ChunkMeta) ([]*dump.Chunk, error) {
	if meta.Index < 0 || meta.Index >= len(s.config.Databases) {
		return nil, fmt.Errorf("postgres chunk index %d out of range", meta.Index)
	}
	dbName := s.config.Databases[meta.Index]

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
		return nil, fmt.Errorf("pg_dump failed: %s: %w", stderr.String(), err)
	}

	return []*dump.Chunk{{
		ChunkMeta: meta,
		Content:   out.Bytes(),
		Filename:  fmt.Sprintf("%d.pgdump", meta.Index),
	}}, nil
}

// WriteChunk restores the PostgreSQL dump using pg_restore.
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
		return fmt.Errorf("pg_restore failed: %s: %w", stderr.String(), err)
	}
	return nil
}

// FinalizeWrites is a no-op for PostgreSQL.
func (s Source) FinalizeWrites() error {
	return nil
}
