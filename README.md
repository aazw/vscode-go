# Go Playgrounds for me

## 概要

* [playgrounds/custom_error_with_stacktrace/](./playgrounds/custom_error_with_stacktrace/)
    * カスタムエラーの実装例
    * StackTrace付きでの取り回しかた
    * slogでのStackTrace付きでのログの履き方

* [playgrounds/errors_stdpkg](./playgrounds/errors_stdpkg)
    * 標準パッケージの1つ『errors』を使った場合、Wrapしたエラーの出力がどうなるかの検証    

* [playgrounds/errors_cockroachdb](./playgrounds/errors_cockroachdb)
    * デファクトライブラリの1つ『cockroachdb/errors』を使った場合、Wrapしたエラーの出力がどうなるかの検証    

* [playgrounds/custom_error_with_stacktrace](./playgrounds/custom_error_with_stacktrace)
    * 独自エラー型の追求
    * エラー発生時のStackTrace
    * エラーハンドリングの経路記録
    * メタデータの扱い
    * など

* [playgrounds/cobra_and_viper](./playgrounds/cobra_and_viper)
    * urfave/cli 以外のcliライブラリの検証
    * cobraを検証してたら、configライブラリのviperもついてきたので一緒に検証
