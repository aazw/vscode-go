# Go Playgrounds for me

## 概要 (昇順)

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

* [playgrounds/go-playground_validator_with_cerrors](./playgrounds/go-playground_validator_with_cerrors)
    * `github.com/go-playground/validator/v10`を使う方法検証
    * 複数エラーの処理方法
    * エラーとなった内容をどう返すかのメッセージ組み立ての検証

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
