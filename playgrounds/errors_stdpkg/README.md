# playgrounds/errors_stdpkg

<https://pkg.go.dev/errors>

## 例

```bash
$ go run -v main.go 
command-line-arguments
err: func001: func002: new on func003

err: func001: func002: new on func003
```

Go 1.13 以降の errors パッケージは以下を提供する.

* 「ラップ機構」（%w、errors.Unwrap/Is/As/Join）――エラーを“つなぐ”ための API
* 「フォーマット機構」（fmt.Formatter を error 型が実装している場合に %+v などへ出力）――エラーを“どう見せるか”を型側に委ねる仕組み

以下は errors パッケージは提供しない.

* 「スタックトレースを自動で付けて %+v で詳細を吐く」機能

より詳細／構造的な出力を得たい場合は次の 3 択になる.

1. 標準 API だけで“構造的”に列挙する (for文でerrorを処理)
2. 自分でスタックトレースを載せるカスタムエラー型を作る
3. 既成ライブラリを使う
