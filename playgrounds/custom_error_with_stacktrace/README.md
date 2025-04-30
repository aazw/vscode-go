## 概要

* カスタムエラーの実装例
* StackTrace付きでの取り回しかた
* slogでのStackTrace付きでのログの履き方

## 例

```bash
$ go run main.go | jq
{
  "time": "2025-04-30T15:18:15.019087433Z",
  "level": "ERROR",
  "source": {
    "function": "main.main",
    "file": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
    "line": 22
  },
  "msg": "error message",
  "err": {
    "code": "UNKNOWN_ERROR",
    "description": "an unknown error occurred",
    "message": "hello world",
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
          "lineno": 22,
          "in_app": true
        },
        {
          "function": "func001",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 26,
          "in_app": true
        },
        {
          "function": "func002",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 30,
          "in_app": true
        },
        {
          "function": "func003",
          "module": "main",
          "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
          "lineno": 34,
          "in_app": true
        }
      ]
    }
  }
}
```
