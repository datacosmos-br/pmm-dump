# Triagem SonarCloud — datacosmos-br/pmm-dump

Gerado do dump da plataforma SonarCloud (2026-08-06).

Bead de rastreio: `ai-hub-31k0.5`

## Resumo

**58 issues** — BLOCKER 2, CRITICAL 31, MAJOR 11, MINOR 9
Tipos: VULNERABILITY 12, BUG 1, CODE_SMELL 45

| regra | issues |
|---|---|
| `go:S3776` | 28 |
| `godre:S8184` | 7 |
| `go:S1135` | 4 |
| `githubactions:S8233` | 3 |
| `go:S6437` | 2 |
| `go:S2068` | 2 |
| `go:S2077` | 2 |
| `go:S4830` | 1 |

## Issues

Coluna **Decisão**: `corrigir` / `falso-positivo` / `risco-aceito`.

| # | sev | tipo | regra | componente | linha | Decisão |
|---|---|---|---|---|---|---|
| 1 | BLOCKER | VULNERABILITY | `go:S6437` | `internal/test/deployment/pmm.go` | 152 | |
| 2 | BLOCKER | VULNERABILITY | `go:S6437` | `internal/test/deployment/pmm.go` | 175 | |
| 3 | CRITICAL | CODE_SMELL | `go:S3776` | `cmd/pmm-dump/main.go` | 45 | |
| 4 | CRITICAL | CODE_SMELL | `go:S3776` | `cmd/pmm-dump/util.go` | 225 | |
| 5 | CRITICAL | CODE_SMELL | `go:S3776` | `cmd/pmm-dump/util.go` | 402 | |
| 6 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/deployment/image.go` | 73 | |
| 7 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/deployment/pmm.go` | 244 | |
| 8 | CRITICAL | VULNERABILITY | `go:S4830` | `internal/test/deployment/pmm.go` | 476 | |
| 9 | CRITICAL | VULNERABILITY | `go:S5527` | `internal/test/deployment/pmm.go` | 476 | |
| 10 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/errmsg_test.go` | 30 | |
| 11 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/qan_test.go` | 48 | |
| 12 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/qan_test.go` | 173 | |
| 13 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/qan_test.go` | 205 | |
| 14 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/qan_test.go` | 257 | |
| 15 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/validate_test.go` | 133 | |
| 16 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/validate_test.go` | 209 | |
| 17 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/validate_test.go` | 261 | |
| 18 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/validate_test.go` | 313 | |
| 19 | CRITICAL | CODE_SMELL | `go:S3776` | `internal/test/e2e/version_test.go` | 32 | |
| 20 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/clickhouse/tsv/tsv.go` | 91 | |
| 21 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/grafana/expr/vm.go` | 127 | |
| 22 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/grafana/expr/vm.go` | 174 | |
| 23 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/grafana/expr/vmquery.go` | 35 | |
| 24 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/grafana/expr/vmtmpl.go` | 33 | |
| 25 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/grafana/templating/interpolate.go` | 35 | |
| 26 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/grafana/templating/interpolate.go` | 114 | |
| 27 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/transferer/export.go` | 138 | |
| 28 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/transferer/export_test.go` | 26 | |
| 29 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/transferer/import.go` | 32 | |
| 30 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/transferer/import_test.go` | 32 | |
| 31 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/transferer/import_test.go` | 155 | |
| 32 | CRITICAL | CODE_SMELL | `go:S1192` | `pkg/victoriametrics/source.go` | 102 | |
| 33 | CRITICAL | CODE_SMELL | `go:S3776` | `pkg/victoriametrics/source_test.go` | 34 | |
| 34 | MAJOR | VULNERABILITY | `githubactions:S8233` | `.github/workflows/ci.yml` | 13 | |
| 35 | MAJOR | VULNERABILITY | `githubactions:S8233` | `.github/workflows/ci.yml` | 14 | |
| 36 | MAJOR | VULNERABILITY | `githubactions:S8233` | `.github/workflows/ci.yml` | 15 | |
| 37 | MAJOR | VULNERABILITY | `githubactions:S7636` | `.github/workflows/ci.yml` | 115 | |
| 38 | MAJOR | VULNERABILITY | `go:S2068` | `internal/test/deployment/container.go` | 173 | |
| 39 | MAJOR | VULNERABILITY | `go:S2068` | `internal/test/deployment/container.go` | 229 | |
| 40 | MAJOR | CODE_SMELL | `go:S107` | `internal/test/deployment/container.go` | 288 | |
| 41 | MAJOR | BUG | `godre:S8168` | `pkg/clickhouse/source.go` | 70 | |
| 42 | MAJOR | VULNERABILITY | `go:S2077` | `pkg/clickhouse/source.go` | 112 | |
| 43 | MAJOR | VULNERABILITY | `go:S2077` | `pkg/clickhouse/source.go` | 218 | |
| 44 | MAJOR | CODE_SMELL | `shelldre:S7688` | `setup/test/init-test-configs.sh` | 11 | |
| 45 | MINOR | CODE_SMELL | `godre:S8193` | `pkg/grafana/types/types.go` | 71 | |
| 46 | MINOR | CODE_SMELL | `godre:S8196` | `pkg/transferer/transferer.go` | 53 | |
| 47 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 21 | |
| 48 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 22 | |
| 49 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 23 | |
| 50 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 24 | |
| 51 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 25 | |
| 52 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 26 | |
| 53 | MINOR | CODE_SMELL | `godre:S8184` | `tools/tools.go` | 27 | |
| 54 | INFO | CODE_SMELL | `githubactions:S1135` | `.github/workflows/ci.yml` | 51 | |
| 55 | INFO | CODE_SMELL | `go:S1135` | `cmd/pmm-dump/main.go` | 257 | |
| 56 | INFO | CODE_SMELL | `go:S1135` | `internal/test/deployment/env.go` | 82 | |
| 57 | INFO | CODE_SMELL | `go:S1135` | `pkg/grafana/templating/interpolate.go` | 157 | |
| 58 | INFO | CODE_SMELL | `go:S1135` | `pkg/grafana/types/types.go` | 19 | |

## Como triar

1. **BLOCKER e CRITICAL primeiro**, e todo VULNERABILITY independente de severidade.
2. Classificar: **corrigir**, **falso-positivo** (marcar na plataforma SonarCloud com justificativa), **risco-aceito** (com prazo).
3. CODE_SMELL em volume alto sugere padrão — corrigir a causa raiz, não issue a issue.

Dados brutos: `~/sonarqube-violations/by-repo/datacosmos-br__pmm-dump.json`

