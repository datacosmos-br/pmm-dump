# Triagem SonarCloud — datacosmos-br/pmm-dump

Gerado do dump da plataforma SonarCloud (2026-08-06).

Bead: `ai-hub-31k0.5`

## Resumo

**58 issues** — BLOCKER 2, CRITICAL 31, MAJOR 11, MINOR 9
Tipos: VULNERABILITY 12, BUG 1, CODE_SMELL 45 · **Debt total: 832min**

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
| `go:S5527` | 1 |
| `go:S1192` | 1 |

## Como usar

Cada issue traz a **mensagem do SonarQube** (descreve o problema e o impacto), o **código real** (linha `>>>`), o tipo e o effort estimado.
**Decisão**: `corrigir` / `falso-positivo` (marcar na plataforma com justificativa) / `risco-aceito`. Ordem: BLOCKER → CRITICAL → VULNERABILITY → MAJOR. CODE_SMELL em volume pede correção de padrão.

## Issues

### 1 · 🔴 BLOCKER · VULNERABILITY · `go:S6437`
**Local**: `internal/test/deployment/pmm.go:152` · **Effort**: 1h

> Revoke and change this secret, as it might be compromised.

```go
      148  	if err != nil {
      149  		p.t.Fatal(err)
      150  	}
      151  	if u.User.Username() == "" {
>>>   152  		u.User = url.UserPassword("admin", "admin")
      153  	}
      154  	if u.Scheme == "" {
      155  		u.Scheme = "http"
      156  	}
```

**Decisão**: 

### 2 · 🔴 BLOCKER · VULNERABILITY · `go:S6437`
**Local**: `internal/test/deployment/pmm.go:175` · **Effort**: 1h

> Revoke and change this secret, as it might be compromised.

```go
      171  	if err != nil {
      172  		p.t.Fatal(err)
      173  	}
      174  
>>>   175  	u.User = url.UserPassword("default", "clickhouse")
      176  
      177  	u.Scheme = "clickhouse"
      178  	u.Path = "pmm"
      179  	if strings.Contains(u.Host, ":") {
```

**Decisão**: 

### 3 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `cmd/pmm-dump/main.go:45` · **Effort**: 2h26min

> Refactor this method to reduce its Cognitive Complexity from 156 to the 15 allowed.

```go
       41  	GitCommit  string
       42  	GitVersion string
       43  )
       44  
>>>    45  func main() { //nolint:gocyclo,maintidx
       46  	var (
       47  		cli = kingpin.New("pmm-dump", "Percona PMM Dump")
       48  
       49  		// general options
```

**Decisão**: 

### 4 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `cmd/pmm-dump/util.go:225` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
      221  	}
      222  	return resp.Timezone, nil
      223  }
      224  
>>>   225  func composeMeta(pmmURL string, c *client.Client, exportServices bool, cli *kingpin.Application, vmNativeData bool) (*dump.Meta, error) {
      226  	_, pmmVer, err := getPMMVersion(pmmURL, c)
      227  	if err != nil {
      228  		return nil, errors.Wrap(err, "failed to get PMM version")
      229  	}
```

**Decisão**: 

### 5 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `cmd/pmm-dump/util.go:402` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
      398  
      399  	return clickhouseSource, true
      400  }
      401  
>>>   402  func parseURL(pmmURL, pmmHost, pmmPort, pmmUser, pmmPassword *string) {
      403  	parsedURL, err := url.Parse(*pmmURL)
      404  	if err != nil {
      405  		log.Fatal().Err(err).Msg("Cannot parse pmm url")
      406  	}
```

**Decisão**: 

### 6 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/deployment/image.go:73` · **Effort**: 9min

> Refactor this method to reduce its Cognitive Complexity from 19 to the 15 allowed.

```go
       69  
       70  	return false, nil
       71  }
       72  
>>>    73  func PullNecessaryImages(ctx context.Context) error {
       74  	files, err := os.ReadDir(util.TestDir)
       75  	if err != nil {
       76  		return errors.Wrap(err, "failed to read test dir")
       77  	}
```

**Decisão**: 

### 7 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/deployment/pmm.go:244` · **Effort**: 11min

> Refactor this method to reduce its Cognitive Complexity from 21 to the 15 allowed.

```go
      240  }
      241  
      242  var checkImagesMu sync.Mutex
      243  
>>>   244  func (pmm *PMM) deploy(ctx context.Context) error {
      245  	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
      246  	if err != nil {
      247  		return errors.Wrap(err, "failed to create docker client")
      248  	}
```

**Decisão**: 

### 8 · 🟠 CRITICAL · VULNERABILITY · `go:S4830`
**Local**: `internal/test/deployment/pmm.go:476` · **Effort**: 5min

> Enable server certificate validation on this SSL/TLS connection.

```go
      472  		MaxIdemponentCallAttempts: 5, //nolint:mnd
      473  		ReadTimeout:               time.Minute,
      474  		WriteTimeout:              time.Minute,
      475  		MaxConnWaitTimeout:        time.Second * 30, //nolint:mnd
>>>   476  		TLSConfig: &tls.Config{
      477  			InsecureSkipVerify: true, //nolint:gosec
      478  		},
      479  	}
      480  	authParams := grafanaClient.AuthParams{
```

**Decisão**: 

### 9 · 🟠 CRITICAL · VULNERABILITY · `go:S5527`
**Local**: `internal/test/deployment/pmm.go:476` · **Effort**: 5min

> Enable server hostname verification on this SSL/TLS connection.

```go
      472  		MaxIdemponentCallAttempts: 5, //nolint:mnd
      473  		ReadTimeout:               time.Minute,
      474  		WriteTimeout:              time.Minute,
      475  		MaxConnWaitTimeout:        time.Second * 30, //nolint:mnd
>>>   476  		TLSConfig: &tls.Config{
      477  			InsecureSkipVerify: true, //nolint:gosec
      478  		},
      479  	}
      480  	authParams := grafanaClient.AuthParams{
```

**Decisão**: 

### 10 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/errmsg_test.go:30` · **Effort**: 10min

> Refactor this method to reduce its Cognitive Complexity from 20 to the 15 allowed.

```go
       26  	"pmm-dump/internal/test/util"
       27  	"pmm-dump/pkg/grafana/client"
       28  )
       29  
>>>    30  func TestErrMsgCheckCompatibilityVersion(t *testing.T) {
       31  	var b util.Binary
       32  	tests := []struct {
       33  		name          string
       34  		internalError bool
```

**Decisão**: 

### 11 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/qan_test.go:48` · **Effort**: 32min

> Refactor this method to reduce its Cognitive Complexity from 42 to the 15 allowed.

```go
       44  const qanTestRetryTimeout = time.Minute * 2
       45  
       46  var qanPMM = deployment.NewReusablePMM("qan", ".env.test")
       47  
>>>    48  func TestQANWhere(t *testing.T) {
       49  	ctx := context.Background()
       50  	c := deployment.NewController(t)
       51  	pmm := c.ReusablePMM(qanPMM)
       52  	if err := pmm.Deploy(ctx); err != nil {
```

**Decisão**: 

### 12 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/qan_test.go:173` · **Effort**: 15min

> Refactor this method to reduce its Cognitive Complexity from 25 to the 15 allowed.

```go
      169  		})
      170  	}
      171  }
      172  
>>>   173  func validateQAN(data []byte, columnTypes []*sql.ColumnType, equalMap map[string]string) error {
      174  	tr := tsv.NewReader(bytes.NewReader(data), columnTypes)
      175  	for {
      176  		values, err := tr.Read()
      177  		if err != nil {
```

**Decisão**: 

### 13 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/qan_test.go:205` · **Effort**: 9min

> Refactor this method to reduce its Cognitive Complexity from 19 to the 15 allowed.

```go
      201  	}
      202  	return nil
      203  }
      204  
>>>   205  func getQANChunks(filename string) (map[string][]byte, error) {
      206  	f, err := os.Open(filename) //nolint:gosec
      207  	if err != nil {
      208  		return nil, err
      209  	}
```

**Decisão**: 

### 14 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/qan_test.go:257` · **Effort**: 11min

> Refactor this method to reduce its Cognitive Complexity from 21 to the 15 allowed.

```go
      253  	}
      254  	return chunkMap, nil
      255  }
      256  
>>>   257  func TestQANEmptyChunks(t *testing.T) {
      258  	ctx := context.Background()
      259  
      260  	c := deployment.NewController(t)
      261  	pmm := c.ReusablePMM(qanPMM)
```

**Decisão**: 

### 15 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/validate_test.go:133` · **Effort**: 24min

> Refactor this method to reduce its Cognitive Complexity from 34 to the 15 allowed.

```go
      129  	pmm.Log(fmt.Sprintf("Data loss in similar chunks is %f%%", loss*100))
      130  	pmm.Log(fmt.Sprintf("Amount of missing chunks is %d", missingChunks))
      131  }
      132  
>>>   133  func validateChunks(t *testing.T, pmm *deployment.PMM, xDump, yDump string) (float64, int, error) {
      134  	t.Helper()
      135  
      136  	xChunkMap, err := readChunks(xDump)
      137  	if err != nil {
```

**Decisão**: 

### 16 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/validate_test.go:209` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
      205  	}
      206  	return float64(totalMissingValues) / float64(totalValues), len(xMissingChunks) + len(yMissingChunks), nil
      207  }
      208  
>>>   209  func chCompareChunks(t *testing.T, pmm *deployment.PMM, filename string, xDump, yDump string, xChunkData, yChunkData []byte) {
      210  	t.Helper()
      211  
      212  	getHashMap := func(data []byte) map[string][]string {
      213  		r := tsv.NewReader(bytes.NewBuffer(data), nil)
```

**Decisão**: 

### 17 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/validate_test.go:261` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
      257  	}
      258  	return total
      259  }
      260  
>>>   261  func vmCompareChunkData(pmm *deployment.PMM, xChunk, yChunk []vmMetric) (int, error) {
      262  	if len(xChunk) != len(yChunk) {
      263  		pmm.Log(fmt.Sprintf("Size of chunks is different: len(x)=%d, len(y)=%d", len(xChunk), len(yChunk)))
      264  	}
      265  
```

**Decisão**: 

### 18 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/validate_test.go:313` · **Effort**: 7min

> Refactor this method to reduce its Cognitive Complexity from 17 to the 15 allowed.

```go
      309  }
      310  
      311  type chunkMap map[string][]byte
      312  
>>>   313  func readChunks(filename string) (chunkMap, error) {
      314  	f, err := os.Open(filename) //nolint:gosec
      315  	if err != nil {
      316  		return nil, err
      317  	}
```

**Decisão**: 

### 19 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `internal/test/e2e/version_test.go:32` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
       28  	"pmm-dump/internal/test/deployment"
       29  	"pmm-dump/internal/test/util"
       30  )
       31  
>>>    32  func TestPMMCompatibility(t *testing.T) {
       33  	ctx := context.Background()
       34  
       35  	pmmVersions, err := getVersions()
       36  	if err != nil {
```

**Decisão**: 

### 20 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/clickhouse/tsv/tsv.go:91` · **Effort**: 18min

> Refactor this method to reduce its Cognitive Complexity from 28 to the 15 allowed.

```go
       87  	}
       88  	return result, nil
       89  }
       90  
>>>    91  func parseElement(record string, st reflect.Type) (interface{}, error) {
       92  	var value interface{}
       93  	var err error
       94  	switch st.Kind() {
       95  	case reflect.Slice:
```

**Decisão**: 

### 21 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/grafana/expr/vm.go:127` · **Effort**: 7min

> Refactor this method to reduce its Cognitive Complexity from 17 to the 15 allowed.

```go
      123  }
      124  
      125  const VMDatasourceName = "Metrics"
      126  
>>>   127  func (p *VMExprParser) GetSelectors(dashboard types.DashboardPanel) ([]string, error) {
      128  	selectorMap := make(map[string]struct{})
      129  
      130  	err := p.parseTemplatingVars(dashboard.Templating.List)
      131  	if err != nil {
```

**Decisão**: 

### 22 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/grafana/expr/vm.go:174` · **Effort**: 13min

> Refactor this method to reduce its Cognitive Complexity from 23 to the 15 allowed.

```go
      170  }
      171  
      172  var errShouldIgnoreQuery = errors.New("should ignore query")
      173  
>>>   174  func (p *VMExprParser) parseTemplatingVar(v types.VariableModel) (templating.TemplatingVariable, error) {
      175  	switch v.Type {
      176  	case types.VariableTypeQuery:
      177  		if v.Datasource != nil {
      178  			uid, err := templating.InterpolateQuery(v.Datasource.UID, p.from, p.to, p.allVariables())
```

**Decisão**: 

### 23 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/grafana/expr/vmquery.go:35` · **Effort**: 16min

> Refactor this method to reduce its Cognitive Complexity from 26 to the 15 allowed.

```go
       31  	value = strings.ReplaceAll(value, `"`, `\"`)
       32  	return value
       33  }
       34  
>>>    35  func (p *VMExprParser) parseQuery(query string) ([]string, error) {
       36  	if query == "" {
       37  		return nil, nil
       38  	}
       39  
```

**Decisão**: 

### 24 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/grafana/expr/vmtmpl.go:33` · **Effort**: 9min

> Refactor this method to reduce its Cognitive Complexity from 19 to the 15 allowed.

```go
       29  	"pmm-dump/pkg/grafana/types"
       30  	"pmm-dump/pkg/victoriametrics"
       31  )
       32  
>>>    33  func (p *VMExprParser) parseTemplatingQuery(v types.VariableModel) (templating.TemplatingVariable, error) {
       34  	query, err := templating.GetQueryFromModel(v)
       35  	if err != nil {
       36  		return templating.TemplatingVariable{}, errors.Wrap(err, "get query from model")
       37  	}
```

**Decisão**: 

### 25 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/grafana/templating/interpolate.go:35` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
       31  	}
       32  	return "", errors.Errorf("unsupported format by pmm-dump: %s", format)
       33  }
       34  
>>>    35  func InterpolateQuery(query string, from time.Time, to time.Time, vars []TemplatingVariable) (string, error) {
       36  	if query == "" {
       37  		return "", nil
       38  	}
       39  	query, err := ApplyMacros(query, TimeRange{From: from, To: to})
```

**Decisão**: 

### 26 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/grafana/templating/interpolate.go:114` · **Effort**: 10min

> Refactor this method to reduce its Cognitive Complexity from 20 to the 15 allowed.

```go
      110  	}
      111  	return TemplatingVariable{}, false
      112  }
      113  
>>>   114  func (v TemplatingVariable) Interpolate(format VariableFormat) (string, error) {
      115  	if format == "" {
      116  		format = FormatPipe
      117  	}
      118  	values := v.Values
```

**Decisão**: 

### 27 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/transferer/export.go:138` · **Effort**: 6min

> Refactor this method to reduce its Cognitive Complexity from 16 to the 15 allowed.

```go
      134  		}
      135  	}
      136  }
      137  
>>>   138  func (t Transferer) writeChunksToFile(meta dump.Meta, chunkC <-chan *dump.Chunk, logBuffer *bytes.Buffer) error {
      139  	gzw, err := gzip.NewWriterLevel(t.file, gzip.BestCompression)
      140  	if err != nil {
      141  		return errors.Wrap(err, "failed to create gzip writer")
      142  	}
```

**Decisão**: 

### 28 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/transferer/export_test.go:26` · **Effort**: 20min

> Refactor this method to reduce its Cognitive Complexity from 30 to the 15 allowed.

```go
       22  
       23  	"pmm-dump/pkg/dump"
       24  )
       25  
>>>    26  func TestExport(t *testing.T) {
       27  	ctx := context.Background()
       28  
       29  	type lsOpts struct {
       30  		status          LoadStatus
```

**Decisão**: 

### 29 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/transferer/import.go:32` · **Effort**: 19min

> Refactor this method to reduce its Cognitive Complexity from 29 to the 15 allowed.

```go
       28  
       29  	"pmm-dump/pkg/dump"
       30  )
       31  
>>>    32  func (t Transferer) Import(ctx context.Context, runtimeMeta dump.Meta) error {
       33  	log.Info().Msg("Importing metrics...")
       34  	gzr, err := gzip.NewReader(t.file)
       35  	if err != nil {
       36  		return errors.Wrap(err, "failed to open as gzip")
```

**Decisão**: 

### 30 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/transferer/import_test.go:32` · **Effort**: 15min

> Refactor this method to reduce its Cognitive Complexity from 25 to the 15 allowed.

```go
       28  
       29  	"pmm-dump/pkg/dump"
       30  )
       31  
>>>    32  func TestImport(t *testing.T) {
       33  	ctx := context.Background()
       34  
       35  	tests := []struct {
       36  		name          string
```

**Decisão**: 

### 31 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/transferer/import_test.go:155` · **Effort**: 33min

> Refactor this method to reduce its Cognitive Complexity from 43 to the 15 allowed.

```go
      151  	writeFakeFile(t, &buf, opts)
      152  	return buf.Bytes()
      153  }
      154  
>>>   155  func writeFakeFile(t *testing.T, w io.Writer, opts fakeFileOpts) {
      156  	t.Helper()
      157  
      158  	gzw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
      159  	if err != nil {
```

**Decisão**: 

### 32 · 🟠 CRITICAL · CODE_SMELL · `go:S1192`
**Local**: `pkg/victoriametrics/source.go:102` · **Effort**: 6min

> Define a constant instead of duplicating this literal "Got successful response from Victoria Metrics" 3 times.

```go
       98  	if status := resp.StatusCode(); status != fasthttp.StatusOK {
       99  		return nil, errors.Errorf("non-OK response from victoria metrics: %d: %s", status, gzipDecode(body))
      100  	}
      101  
>>>   102  	log.Debug().Msg("Got successful response from Victoria Metrics")
      103  
      104  	chunk := &dump.Chunk{
      105  		ChunkMeta: m,
      106  		Content:   body,
```

**Decisão**: 

### 33 · 🟠 CRITICAL · CODE_SMELL · `go:S3776`
**Local**: `pkg/victoriametrics/source_test.go:34` · **Effort**: 18min

> Refactor this method to reduce its Cognitive Complexity from 28 to the 15 allowed.

```go
       30  
       31  	"pmm-dump/pkg/grafana/client"
       32  )
       33  
>>>    34  func TestWriteChunk(t *testing.T) {
       35  	tests := []struct {
       36  		name         string
       37  		metricsSize  int
       38  		contentLimit int
```

**Decisão**: 

### 34 · 🟡 MAJOR · VULNERABILITY · `githubactions:S8233`
**Local**: `.github/workflows/ci.yml:13` · **Effort**: 5min

> Move this write permission from workflow level to job level.

```yaml
        9    pull_request:
       10  
       11  permissions:
       12    contents: read
>>>    13    packages: write
       14    checks: write
       15    pull-requests: write
       16  
       17  jobs:
```

**Decisão**: 

### 35 · 🟡 MAJOR · VULNERABILITY · `githubactions:S8233`
**Local**: `.github/workflows/ci.yml:14` · **Effort**: 5min

> Move this write permission from workflow level to job level.

```yaml
       10  
       11  permissions:
       12    contents: read
       13    packages: write
>>>    14    checks: write
       15    pull-requests: write
       16  
       17  jobs:
       18    test:
```

**Decisão**: 

### 36 · 🟡 MAJOR · VULNERABILITY · `githubactions:S8233`
**Local**: `.github/workflows/ci.yml:15` · **Effort**: 5min

> Move this write permission from workflow level to job level.

```yaml
       11  permissions:
       12    contents: read
       13    packages: write
       14    checks: write
>>>    15    pull-requests: write
       16  
       17  jobs:
       18    test:
       19      if: github.event_name == 'pull_request'
```

**Decisão**: 

### 37 · 🟡 MAJOR · VULNERABILITY · `githubactions:S7636`
**Local**: `.github/workflows/ci.yml:115` · **Effort**: 15min

> Avoid expanding secrets in a run block.

```yaml
      111  
      112        - name: Run checks/linters
      113          run: |
      114            # use GITHUB_TOKEN because only it has access to GitHub Checks API
>>>   115            bin/golangci-lint run --out-format=line-number | env REVIEWDOG_GITHUB_API_TOKEN=${{ secrets.GITHUB_TOKEN }} bin/reviewdog -f=golangci-lint -reporter=github-pr-review -filter-mode=added -fail-level=error
      116  
      117            # run it like that until some of those issues/PRs are resolved:
      118            # * https://github.com/quasilyte/go-consistent/issues/33
      119            # * https://github.com/golangci/golangci-lint/issues/288
```

**Decisão**: 

### 38 · 🟡 MAJOR · VULNERABILITY · `go:S2068`
**Local**: `internal/test/deployment/container.go:173` · **Effort**: 30min

> "PASSWORD" detected here, make sure this is not a hard-coded credential.

```go
      169  func (pmm *PMM) CreatePMMClient(ctx context.Context, dockerCli *client.Client, networkID string) error {
      170  	envs := []string{
      171  		"PMM_AGENT_CONFIG_FILE=config/pmm-agent.yaml",
      172  		"PMM_AGENT_SERVER_USERNAME=admin",
>>>   173  		"PMM_AGENT_SERVER_PASSWORD=admin",
      174  		"PMM_AGENT_SERVER_ADDRESS=" + pmm.ServerContainerName() + ":8443",
      175  		"PMM_AGENT_SERVER_INSECURE_TLS=1",
      176  		"PMM_AGENT_SETUP=1",
      177  		"PMM_AGENT_SETUP_FORCE=true",
```

**Decisão**: 

### 39 · 🟡 MAJOR · VULNERABILITY · `go:S2068`
**Local**: `internal/test/deployment/container.go:229` · **Effort**: 30min

> "PASSWORD" detected here, make sure this is not a hard-coded credential.

```go
      225  
      226  	envs := []string{
      227  		"MONGO_INITDB_DATABASE=admin",
      228  		"MONGO_INITDB_ROOT_USERNAME=admin",
>>>   229  		"MONGO_INITDB_ROOT_PASSWORD=admin",
      230  	}
      231  	vol, err := dockerCli.VolumeCreate(ctx, volume.CreateOptions{
      232  		Name: pmm.MongoContainerName() + volumeSuffix,
      233  		Labels: map[string]string{
```

**Decisão**: 

### 40 · 🟡 MAJOR · CODE_SMELL · `go:S107`
**Local**: `internal/test/deployment/container.go:288` · **Effort**: 20min

> This function has 10 parameters, which is greater than the 7 authorized.

```go
      284  
      285  	return nil
      286  }
      287  
>>>   288  func (pmm *PMM) createContainer(ctx context.Context,
      289  	dockerCli *client.Client,
      290  	name,
      291  	image string,
      292  	ports []string,
```

**Decisão**: 

### 41 · 🟡 MAJOR · BUG · `godre:S8168`
**Local**: `pkg/clickhouse/source.go:70` · **Effort**: 5min

> Add 'defer tx.Rollback()' after checking the error from 'db.Begin()' to ensure the transaction is rolled back on failure.

```go
       66  	}
       67  
       68  	log.Debug().Str("engine", engine).Str("table", tableName).Msg("Detected ClickHouse table")
       69  
>>>    70  	tx, err := db.Begin()
       71  	if err != nil {
       72  		return nil, errors.Wrap(err, "begin")
       73  	}
       74  
```

**Decisão**: 

### 42 · 🟡 MAJOR · VULNERABILITY · `go:S2077`
**Local**: `pkg/clickhouse/source.go:112` · **Effort**: 20min

> Make sure using a dynamically formatted SQL query is safe here.

```go
      108  	return engine, tableName, nil
      109  }
      110  
      111  func columnTypes(db *sql.DB, tableName string) ([]*sql.ColumnType, error) {
>>>   112  	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 1", tableName))
      113  	if err != nil {
      114  		return nil, err
      115  	}
      116  	defer rows.Close() //nolint:errcheck
```

**Decisão**: 

### 43 · 🟡 MAJOR · VULNERABILITY · `go:S2077`
**Local**: `pkg/clickhouse/source.go:218` · **Effort**: 20min

> Make sure using a dynamically formatted SQL query is safe here.

```go
      214  	for i := 0; i < columnsCount-1; i++ {
      215  		query.WriteString("?,")
      216  	}
      217  	query.WriteString("?)")
>>>   218  	return tx.Prepare(query.String())
      219  }
      220  
      221  func (s Source) FinalizeWrites() error {
      222  	if err := s.stmt.Close(); err != nil {
```

**Decisão**: 

### 44 · 🟡 MAJOR · CODE_SMELL · `shelldre:S7688`
**Local**: `setup/test/init-test-configs.sh:11` · **Effort**: 2min

> Use '[[' instead of '[' for conditional tests. The '[[' construct is safer and more feature-rich.

```bash
        7  	shift 1
        8  	local append_vars=("${@}")
        9  
       10  	# Return if env file exists
>>>    11  	if [ -e "$target_file" ]; then
       12  		return
       13  	fi
       14  
       15  	printf "# This is a configuration file for running tests.\n# Feel free to modify the values to suit your needs.\n\n" > "$target_file"
```

**Decisão**: 

### 45 · ⚪ MINOR · CODE_SMELL · `godre:S8193`
**Local**: `pkg/grafana/types/types.go:71` · **Effort**: 5min

> Remove this unnecessary variable declaration and use the expression directly in the condition.

```go
       67  		var sort VariableSort
       68  		if err := json.Unmarshal(aux.Sort, &sort); err != nil {
       69  			// Try as int
       70  			var sortInt int
>>>    71  			if err2 := json.Unmarshal(aux.Sort, &sortInt); err2 == nil {
       72  				sort.Type = sortInt
       73  				sort.Desc = false
       74  			} else {
       75  				return err
```

**Decisão**: 

### 46 · ⚪ MINOR · CODE_SMELL · `godre:S8196`
**Local**: `pkg/transferer/transferer.go:53` · **Effort**: 5min

> Rename this interface to follow Go naming conventions for single-method interfaces.

```go
       49  		file:         file,
       50  	}, nil
       51  }
       52  
>>>    53  type ChunkPool interface {
       54  	Next() (dump.ChunkMeta, bool)
       55  }
       56  
       57  type LoadStatusGetter interface {
```

**Decisão**: 

### 47 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:21` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       17  
       18  package tools
       19  
       20  import (
>>>    21  	_ "github.com/apache/skywalking-eyes/cmd/license-eye"
       22  	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
       23  	_ "github.com/quasilyte/go-consistent"
       24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
       25  	_ "golang.org/x/tools/cmd/goimports"
```

**Decisão**: 

### 48 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:22` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       18  package tools
       19  
       20  import (
       21  	_ "github.com/apache/skywalking-eyes/cmd/license-eye"
>>>    22  	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
       23  	_ "github.com/quasilyte/go-consistent"
       24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
       25  	_ "golang.org/x/tools/cmd/goimports"
       26  	_ "gopkg.in/reform.v1/reform"
```

**Decisão**: 

### 49 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:23` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       19  
       20  import (
       21  	_ "github.com/apache/skywalking-eyes/cmd/license-eye"
       22  	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
>>>    23  	_ "github.com/quasilyte/go-consistent"
       24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
       25  	_ "golang.org/x/tools/cmd/goimports"
       26  	_ "gopkg.in/reform.v1/reform"
       27  	_ "mvdan.cc/gofumpt"
```

**Decisão**: 

### 50 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:24` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       20  import (
       21  	_ "github.com/apache/skywalking-eyes/cmd/license-eye"
       22  	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
       23  	_ "github.com/quasilyte/go-consistent"
>>>    24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
       25  	_ "golang.org/x/tools/cmd/goimports"
       26  	_ "gopkg.in/reform.v1/reform"
       27  	_ "mvdan.cc/gofumpt"
       28  )
```

**Decisão**: 

### 51 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:25` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       21  	_ "github.com/apache/skywalking-eyes/cmd/license-eye"
       22  	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
       23  	_ "github.com/quasilyte/go-consistent"
       24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
>>>    25  	_ "golang.org/x/tools/cmd/goimports"
       26  	_ "gopkg.in/reform.v1/reform"
       27  	_ "mvdan.cc/gofumpt"
       28  )
       29  
```

**Decisão**: 

### 52 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:26` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       22  	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
       23  	_ "github.com/quasilyte/go-consistent"
       24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
       25  	_ "golang.org/x/tools/cmd/goimports"
>>>    26  	_ "gopkg.in/reform.v1/reform"
       27  	_ "mvdan.cc/gofumpt"
       28  )
       29  
       30  //go:generate env CGO_ENABLED=0 go build -o ../bin/license-eye github.com/apache/skywalking-eyes/cmd/license-eye
```

**Decisão**: 

### 53 · ⚪ MINOR · CODE_SMELL · `godre:S8184`
**Local**: `tools/tools.go:27` · **Effort**: 5min

> Add a comment explaining why this blank import is needed.

```go
       23  	_ "github.com/quasilyte/go-consistent"
       24  	_ "github.com/reviewdog/reviewdog/cmd/reviewdog"
       25  	_ "golang.org/x/tools/cmd/goimports"
       26  	_ "gopkg.in/reform.v1/reform"
>>>    27  	_ "mvdan.cc/gofumpt"
       28  )
       29  
       30  //go:generate env CGO_ENABLED=0 go build -o ../bin/license-eye github.com/apache/skywalking-eyes/cmd/license-eye
       31  //go:generate go build -o ../bin/go-consistent github.com/quasilyte/go-consistent
```

**Decisão**: 

### 54 · ⚪ INFO · CODE_SMELL · `githubactions:S1135`
**Local**: `.github/workflows/ci.yml:51` · **Effort**: 0min

> Complete the task associated to this "TODO" comment.

```yaml
       47          run: |
       48            go clean -testcache
       49            PMM_DUMP_MAX_PARALLEL_TESTS=3 make run-tests
       50  
>>>    51        #   TODO: Add codecoverage
       52        # - name: Upload coverage results
       53        #   uses: codecov/codecov-action@v3
       54        #   with:
       55        #     token: ${{ secrets.CODECOV_TOKEN }}
```

**Decisão**: 

### 55 · ⚪ INFO · CODE_SMELL · `go:S1135`
**Local**: `cmd/pmm-dump/main.go:257` · **Effort**: 0min

> Complete the task associated to this TODO comment.

```go
      253  		defer file.Close() //nolint:errcheck
      254  
      255  		t, err := transferer.New(file, sources, *workersCount)
      256  		if err != nil {
>>>   257  			log.Fatal().Msgf("Failed to setup export: %v", err) //nolint:gocritic //TODO: potential problem here, see muted linter warning
      258  		}
      259  
      260  		var chunks []dump.ChunkMeta
      261  
```

**Decisão**: 

### 56 · ⚪ INFO · CODE_SMELL · `go:S1135`
**Local**: `internal/test/deployment/env.go:82` · **Effort**: 0min

> Complete the task associated to this TODO comment.

```go
       78  		return "mongo"
       79  	case envVarMongoTag:
       80  		return "latest"
       81  	case envVarPMMVersion:
>>>    82  		// TODO: update this once PMM v3 goes GA
       83  		return "3-dev-latest"
       84  	case envVarUseExistingPMM:
       85  		return "false"
       86  	default:
```

**Decisão**: 

### 57 · ⚪ INFO · CODE_SMELL · `go:S1135`
**Local**: `pkg/grafana/templating/interpolate.go:157` · **Effort**: 0min

> Complete the task associated to this TODO comment.

```go
      153  	if len(values) > 0 {
      154  		return FormatVar(format, values)
      155  	}
      156  
>>>   157  	s, _ := FormatVar(format, values) // TODO: regex escape
      158  	return "(" + s + ")", nil
      159  }
```

**Decisão**: 

### 58 · ⚪ INFO · CODE_SMELL · `go:S1135`
**Local**: `pkg/grafana/types/types.go:19` · **Effort**: 0min

> Complete the task associated to this TODO comment.

```go
       15  package types
       16  
       17  import "encoding/json"
       18  
>>>    19  // TODO: use https://github.com/grafana/grok cli for generating these types
       20  
       21  type DashboardPanel struct {
       22  	Title   string           `json:"title"`
       23  	Panels  []DashboardPanel `json:"panels"`
```

**Decisão**: 

