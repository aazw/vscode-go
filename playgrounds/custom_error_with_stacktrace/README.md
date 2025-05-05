# playgrounds/custom_error_with_stacktrace

* https://github.com/cockroachdb/errors
* https://github.com/getsentry/sentry-go

## 例

```bash
$  go run main.go -log-format json | jq
{
  "time": "2025-05-05T08:42:37.169363861Z",
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
    "cause": "hoge error",
    "stacktrace": [
      {
        "function": "func003",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "lineno": 56,
        "in_app": true
      },
      {
        "function": "func002",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "lineno": 48,
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
        "function": "main",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "lineno": 36,
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
        "function": "goexit",
        "module": "runtime",
        "filename": "runtime/asm_arm64.s",
        "abs_path": "/usr/local/go/src/runtime/asm_arm64.s",
        "lineno": 1223,
        "in_app": true
      }
    ],
    "stacktrace_order": "newest_first",
    "messages": [
      {
        "function": "func003",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "lineno": 56,
        "in_app": false,
        "message": "hello world"
      },
      {
        "function": "func002",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "lineno": 50,
        "in_app": false,
        "message": "huge"
      }
    ],
    "checkpoints": [
      {
        "function": "func001",
        "module": "main",
        "filename": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "abs_path": "/workspaces/vscode-go/playgrounds/custom_error_with_stacktrace/main.go",
        "lineno": 42,
        "in_app": false
      }
    ]
  }
}
```

## Memo (ChatGPT)


```prompt
errors.Is、errors.Asが可能なerrorsパッケージについて、機能差分な星取表をください.

パッケージ例:
* errors
* github.com/pkg/errors
* github.com/cockroachdb/errors
* ほかメジャーなものがあれば必要に応じて随時追加
```

## 最終的なエラーパッケージまとめ

| 記号 | 意味               |
| -- | ---------------- |
| ✓  | 公式にサポート（そのまま使える） |
| △  | 限定的／ラッパー経由なら可    |
| ✗  | 非対応・自前実装が必要      |

| パッケージ                             | Wrap (`%w` / Wrapf) | `errors.Is` / `errors.As` | スタック取得 | `%+v` 等でスタック整形 | エラーコード／メタ情報 |     複数エラー結合     | PII 削除・ネット移送 |
| --------------------------------- | :-----------------: | :-----------------------: | :----: | :------------: | :---------: | :-------------: | :----------: |
| **std errors (Go 1.20+)**         |          ✓          |             ✓             |    ✗   |        ✗       |      ✗      | ✓ `errors.Join` |       ✗      |
| **golang.org/x/xerrors**          |          ✓          |             ✓             |    ✓   |        ✓       |      ✗      |        ✗        |       ✗      |
| **github.com/pkg/errors**         |          ✓          |             ✓             |    ✓   |        ✓       |      ✗      |        ✗        |       ✗      |
| **github.com/cockroachdb/errors** |          ✓          |             ✓             |    ✓   |        ✓       |      ✓      |        ✓        |       ✓      |
| **github.com/rotisserie/eris**    |          ✓          |             ✓             |    ✓   |        ✓       |      ✗      |        ✗        |       ✗      |
| **go.uber.org/multierr**          |       ✗ (集約専用)      |             ✓             |    ✗   |        ✗       |      ✗      |        ✓        |       ✗      |

* PII
  * PII (Personally Identifiable Information) は「個人を識別できる情報」の総称
  * PIIの漏洩リスク対策があるかどうか
    * ログ出力や外部エラートラッキング(Sentry など)へエラーを送るとき、パスワード / API キー / 個人名 / 住所 といった PII をうっかり含むと漏えいになる
    * これに対応する機能があるかどうか

#### 補足・根拠

* **std errors** – Go 1.13 以降 `Is`/`As`/`Unwrap` を実装。Go 1.20 で `errors.Join` が追加され複数エラーを 1 つにまとめられるようになった([Go Packages][1], [Go][2])
* **xerrors** – 標準 `errors` 取り込み前の試験実装で、`Is`/`As` とフレーム付き `xerrors.Errorf` を提供し、`%+v` でスタックを整形表示できる([Go Packages][3])
* **pkg/errors** – `Wrap/WithStack` でスタックを捕捉し、`As`/`Is`/`Unwrap` も補完。メンテナンスモード（非推奨）だが依然広く使われる([Go Packages][4], [GitHub][5], [GitHub][6])
* **cockroachdb/errors** – `pkg/errors` 互換 API に加え、`Cause`/`CombineErrors`/エラーコード/安全詳細/PII 除去（Sentry 連携）といった分散システム向け機能を持つ([Go Packages][7], [Go Packages][8])
* **eris** – スタック付きラップと柔軟なフォーマット、独自の `eris.Is/As` でチェーン照合が可能。コードやマルチエラー機能は持たない([GitHub][9])
* **multierr** – 目的は *結合* に特化。`Combine/Append` で複数エラーをまとめ、戻り値は `errors.Is/As` と相互運用可能([Go Packages][10])

#### 選択の指針（ざっくり）

* **最小依存 & 標準準拠** … `std errors` ＋ `fmt.Errorf("%w")`（Go 1.20 なら `Join` も）
* **既存コードの延命** … `pkg/errors`（ただし新規開発には非推奨）
* **分散システムで安全詳細や互換性が欲しい** … `cockroachdb/errors`
* **可読性重視のスタック表示／JSON 出力** … `eris`
* **単に複数エラーをまとめたいだけ** … `multierr`（または Go 1.20 標準 `Join`）

エラーハンドリング戦略を固める際は、

1. **ラップ箇所を最小化**（同じスタックを何度も付けない）
2. **責務を分ける** – *伝搬* と *ログ出力* でレイヤを分離
3. **Is/As に合わせたカスタム型** – sentinel/error type で比較できるようにする

を意識すると、ライブラリ差異の影響を抑えつつ保守しやすくなります。

[1]: https://pkg.go.dev/errors?utm_source=chatgpt.com "errors - Go Packages"
[2]: https://go.dev/doc/go1.20?utm_source=chatgpt.com "Go 1.20 Release Notes - The Go Programming Language"
[3]: https://pkg.go.dev/golang.org/x/xerrors?utm_source=chatgpt.com "xerrors package - golang.org/x/xerrors - Go Packages"
[4]: https://pkg.go.dev/github.com/pkg/errors "errors package - github.com/pkg/errors - Go Packages"
[5]: https://github.com/moby/moby/discussions/46358?utm_source=chatgpt.com "decide on the future use of `errors`, `pkg/errors` and other errors"
[6]: https://github.com/docker/cli/issues/3618?utm_source=chatgpt.com "Consider dropping dependency on github.com/pkg/errors · Issue ..."
[7]: https://pkg.go.dev/github.com/cockroachdb/errors?utm_source=chatgpt.com "errors package - github.com/cockroachdb/errors - Go Packages"
[8]: https://pkg.go.dev/github.com/cockroachdb/errors "errors package - github.com/cockroachdb/errors - Go Packages"
[9]: https://github.com/rotisserie/eris "GitHub - rotisserie/eris: Error handling library with readable stack traces and flexible formatting support "
[10]: https://pkg.go.dev/go.uber.org/multierr "multierr package - go.uber.org/multierr - Go Packages"

## 複数エラーを結合するユースケース (ChatGPT)

| ユースケース                 | 典型シナリオ                                                   | 何を**結合**するか         | 受け手側のメリット                                   |
| ---------------------- | -------------------------------------------------------- | ------------------- | ------------------------------------------- |
| **① バリデーションの一括報告**     | フォーム入力・設定ファイル・API リクエスト構造体の検証                            | 各フィールドの検証エラー        | ユーザーに「全部まとめて」指摘できる ─― 1 つ直したら次でまた失敗…のループを防ぐ |
| **② 並列処理の収集**          | `errgroup.Group`/`sync.WaitGroup` で並列に I/O・計算・API 呼び出し   | Goroutine ごとの戻り値エラー | 呼び出し元が **1 回の `if err != nil`** で全体結果を把握できる |
| **③ バッチ／一括操作**         | 複数ファイルのアップロード、DB への一括 INSERT、メールの一斉送信                    | 個々のエンティティで発生した失敗    | 何件成功・失敗したかを一覧でロギング・再試行できる                   |
| **④ リソースの解放・後始末**      | `defer` で複数の `Close()` を順次呼ぶ<br>（ファイル, ネット接続, トランザクション等） | 各 `Close` の戻り値      | 後始末の失敗を**取りこぼさず**呼び出し元に返せる（最初のエラーだけ消えない）    |
| **⑤ マルチストレージ／フェイルオーバ** | 同じデータを S3 と GCS に書き込み、どちらかが落ちても良い                        | ストレージごとの失敗原因        | どちらが落ちたかを示しつつ、許容範囲内か判断できる                   |
| **⑥ 複数段階のトランザクション**    | 1 つの API 呼び出し内で「検証→データ登録→通知」のように段階が分かれる                  | 段階ごとの失敗             | 呼び出し元で「どこまで進んだか」を含めてリカバリ判断                  |
