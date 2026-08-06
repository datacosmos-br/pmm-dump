# Triagem Snyk Code (SAST) — datacosmos-br/pmm-dump

Gerado do scan Snyk da org Datacosmos (dump 2026-08-06).

**6 achados** — critical 0, high 1, medium 0, low 5

| categoria | achados |
|---|---|
| Improper Certificate Validation - Permissive TrustManager | 2 |
| SQL Injection | 1 |
| Sensitive Cookie Without 'HttpOnly' Flag | 1 |
| Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | 1 |
| Use of Hardcoded Credentials | 1 |

## Achados

Coluna **Decisão**: `corrigir` / `falso-positivo` / `risco-aceito`.

| # | sev | categoria | arquivo | linha | CWE | Decisão |
|---|---|---|---|---|---|---|
| 1 | high | SQL Injection | `pkg/clickhouse/source.go` | 112 | - | |
| 2 | low | Improper Certificate Validation - Permissive TrustManager | `internal/test/deployment/pmm.go` | 477 | - | |
| 3 | low | Sensitive Cookie Without 'HttpOnly' Flag | `internal/test/e2e/errmsg_test.go` | 77 | - | |
| 4 | low | Sensitive Cookie in HTTPS Session Without 'Secure' Attribute | `internal/test/e2e/errmsg_test.go` | 77 | - | |
| 5 | low | Improper Certificate Validation - Permissive TrustManager | `pkg/victoriametrics/source_test.go` | 71 | - | |
| 6 | low | Use of Hardcoded Credentials | `setup/mongo/init.js` | 21 | - | |

## Como triar

1. Abrir `arquivo:linha` e seguir o fluxo de dados até o sink.
2. Classificar: **corrigir** (entrada externa alcança o sink sem sanitização), **falso-positivo** (credencial de fixture, path de constante — registrar em `.snyk` com justificativa), **risco-aceito** (com prazo de revisão).

Dados brutos: `~/snyk-violations/sast/datacosmos-br__pmm-dump.sast.json`

