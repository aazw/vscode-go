# playgrounds/errors_cockroachdb

<https://github.com/cockroachdb/errors>

## 例

```bash
$ go run -v main.go 
command-line-arguments
err: wrap on func001: wrap on func002: new on func003
(1) attached stack trace
  -- stack trace:
  | main.func001
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:20
  | main.main
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:10
  | runtime.main
  |     /usr/local/go/src/runtime/proc.go:283
Wraps: (2) wrap on func001
Wraps: (3) attached stack trace
  -- stack trace:
  | main.func002
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:28
  | [...repeated from below...]
Wraps: (4) wrap on func002
Wraps: (5) attached stack trace
  -- stack trace:
  | main.func003
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:34
  | main.func002
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:26
  | main.func001
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:18
  | main.main
  |     /workspaces/vscode-go/playgrounds/errors_cockroachdb/main.go:10
  | runtime.main
  |     /usr/local/go/src/runtime/proc.go:283
  | runtime.goexit
  |     /usr/local/go/src/runtime/asm_arm64.s:1223
Wraps: (6) new on func003
Error types: (1) *withstack.withStack (2) *errutil.withPrefix (3) *withstack.withStack (4) *errutil.withPrefix (5) *withstack.withStack (6) *errutil.leafError
```
