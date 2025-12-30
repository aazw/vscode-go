# Go Playgrounds for me

## 概要 (昇順)

* [playgrounds/ast_validation_test_generator](./playgrounds/ast_validation_test_generator)
    * Go AST を使ったバリデーションテストケース自動生成ツール
    * 構造体の `validate` タグ（go-playground/validator形式）を解析
    * 境界値分析・同値分割に基づくテストケースを `go generate` で自動生成
    * 実行方法: `go generate ./playgrounds/ast_validation_test_generator/target/...`
    * 生成ファイル:
        * `{元ファイル名}_{構造体名小文字}_autogen_test.go` - テストケースデータ
        * `unittestgen_helper_autogen_test.go` - ヘルパー型定義
    * 注意: `go run` での直接実行は不可（`go:generate` 経由で設定される環境変数に依存）
    * 未実装: 構造体名を引数で指定する機能（現状は `go:generate` の次の行の構造体のみ対象）

* [playgrounds/cobra_and_viper](./playgrounds/cobra_and_viper)
    * `github.com/urfave/cli` 以外のcliライブラリの検証
    * cobraを検証してたら、configライブラリのviperもついてきたので一緒に検証

* [playgrounds/custom_error_with_stacktrace/](./playgrounds/custom_error_with_stacktrace/)
    * カスタムエラーの実装例
        * 独自エラー型の追求
    * エラー発生時のStackTrace
    * StackTrace付きでのエラーの取り回しかた
    * slogでのStackTrace付きでのログの履き方
    * エラーハンドリングの経路記録
    * メタデータの扱い
    * など

* [playgrounds/errors_cockroachdb](./playgrounds/errors_cockroachdb)
    * デファクトライブラリの1つ `github.com/cockroachdb/errors` を使った場合、Wrapしたエラーの出力がどうなるかの検証    

* [playgrounds/errors_stdpkg](./playgrounds/errors_stdpkg)
    * 標準パッケージの1つ『errors』を使った場合、Wrapしたエラーの出力がどうなるかの検証    

* [playgrounds/generate_id](./playgrounds/generate_id)
    * UUIDv1 〜 UUIDv7までの動作検証
    * xid、ulidも検証

* [playgrounds/gin_graceful_shutdown](./playgrounds/gin_graceful_shutdown)
    * GinでGraceful Shutdownを実装する方法

* [playgrounds/gin_validator](./playgrounds/gin_validator)
    * go-playground/validator
    * そのまま使う場合、Gin組み込みのものをGin経由で使う場合
    * タグの検証

* [playgrounds/gin_with_grafana_stack](./playgrounds/gin_with_grafana_stack)
    * Grafana Stack
        * Grafana、Mimir、Prometheus、Alloy、Loki、Promtail、Tempo、Pyroscope
    * Prometheus Exporter
    * Gin

* [playgrounds/gin_with_valkey](./playgrounds/gin_with_valkey)
    * Redis互換のOSS『Valkey』
    * ginのsession管理にRedis向けのパッケージでValkeyが使えるかの検証

* [playgrounds/go_generate](./playgrounds/go_generate)
    * go generateのお試し

* [playgrounds/go-playground_validator_with_cerrors](./playgrounds/go-playground_validator_with_cerrors)
    * `github.com/go-playground/validator/v10`を使う方法検証
    * 複数エラーの処理方法
    * エラーとなった内容をどう返すかのメッセージ組み立ての検証

* [playgrounds/go_generate](./playgrounds/go_generate)
    * `go geneate`お試し

* [playgrounds/openapi_codegen](./playgrounds/openapi_codegen)
    * OpenAPI Specからコードの生成
    * oapi-codegen

* [playgrounds/scs_with_redis](./playgrounds/scs_with_redis)
    * セッション管理用パッケージお試し

* [playgrounds/sqlc_with_golang_migrate](./playgrounds/sqlc_with_golang_migrate)
    * sqlc、jackc/pgx、golang-migrateの利用方法
    * sqlcでのGoコード生成
    * pgxとsqlcの連携
    * pgxでのコネクションプールの実装
    * golang-migrateでのマイグレーション
    * 上記すべてを考慮したディレクトリ構成

## 外部リンク

* **Cobra**
    * **Cobra**  
      A Framework for Modern CLI Apps in Go  
      https://cobra.dev/
    * https://github.com/spf13/cobra

* **cockroachdb/errors**
    * https://github.com/cockroachdb/errors

* **go-playground/validator**
    * https://github.com/go-playground/validator

* **Visual Studio Code Debugger**
    * **Debug code with Visual Studio Code**  
      https://code.visualstudio.com/docs/debugtest/debugging
    * **Visual Studio Code debug configuration**  
      https://code.visualstudio.com/docs/debugtest/debugging-configuration
