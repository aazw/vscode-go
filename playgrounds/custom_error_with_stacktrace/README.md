## 概要

* カスタムエラーの実装例
* StackTrace付きでの取り回しかた
* slogでのStackTrace付きでのログの履き方

## 例

```bash
$ go run main.go -log-format json | jq
{
  "time": "2025-05-01T05:13:13.044973314Z",
  "level": "ERROR",
  "source": {
    "function": "main.main",
    "file": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
    "line": 36
  },
  "msg": "error message",
  "err": {
    "code": "UNKNOWN_ERROR",
    "detail": "an unknown error occurred",
    "context": [
      {
        "message": "hello world",
        "file": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "line": 48
      },
      {
        "message": "huge",
        "file": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "line": 44
      }
    ],
    "cause": "hoge error",
    "stacktrace": {
      "frames": [
        {
          "function": "goexit",
          "module": "runtime",
          "filename": "runtime/asm_arm64.s",
          "abs_path": "/usr/local/go/src/runtime/asm_arm64.s",
          "lineno": 1223,
          "in_app": true
        },
        {
          "function": "main",
          "module": "runtime",
          "filename": "runtime/proc.go",
          "abs_path": "/usr/local/go/src/runtime/proc.go",
          "lineno": 283,
          "in_app": true
        },
        {
          "function": "main",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 36,
          "in_app": true
        },
        {
          "function": "func001",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 40,
          "in_app": true
        },
        {
          "function": "func002",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 44,
          "in_app": true
        },
        {
          "function": "func003",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 48,
          "in_app": true
        }
      ]
    }
  }
}
```
