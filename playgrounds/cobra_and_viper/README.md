# playgrounds/cobra_cli

* https://github.com/spf13/cobra
* https://cobra.dev/
* https://github.com/urfave/cli
* https://cli.urfave.org/

## cobra使い方 その1: cobra？viper？

https://cobra.dev/#create-rootcmd

> ```go
> func init() {
>   cobra.OnInitialize(initConfig)
>   rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
>   rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
>   rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
>   rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
>   rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
>   viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
>   viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
>   viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
>   viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
>   viper.SetDefault("license", "apache")
> }
> ```

登場人物(?)

* https://github.com/spf13/cobra
* https://github.com/spf13/viper

Cobra で UX の良い CLI を作り、Viper で “設定ソースの一本化” をするのが定番パターン

* Cobra だけだと「設定ファイルや環境変数も使いたい」と思った時に自前でマージ処理を書く必要がある
* Viper だけだと「--help にフラグ一覧を出したい」「サブコマンドごとにフラグを分けたい」等が面倒

|         | **rootCmd.PersistentFlags()（Cobra）**                              | **Viper**                                                                            |
| ------- | ----------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| 役割      | **CLI フラグ定義とパース**<br>– `--flag`, `-f` などのコマンドライン引数を扱う             | **設定ストア**<br>– フラグ・環境変数・設定ファイル（YAML/TOML/JSON…）すべてを 1 つの API で取得できる                  |
| スコープ    | `PersistentFlags()` は**ルート + すべてのサブコマンド**<br>`Flags()` だとそのコマンド限定 | プログラム全体（`viper.Get*()` でどこからでも参照可）                                                   |
| データ取得順序 | コマンドラインに与えられた値だけ                                                  | 優先度を内部で持つ（通常は①明示的 `Set()` ②バインドされたフラグ値 ③環境変数 ④設定ファイル ⑤`SetDefault()`）                |
| 型付け     | フラグ定義時に型が決まる（`StringVar`/`Bool` など）                               | 取り出すときに型指定（`GetString`, `GetBool` …）                                                 |
| ライブラリ依存 | `spf13/cobra` のみ                                                  | `spf13/viper`（内部で `pflag` も利用）                                                       |
| 主な使いどころ | **「CLI で上書きしたい値」**<br>例: `--config`, `--verbose`                  | **「複数ソースをマージした設定」**<br>例: 環境変数 `APP_PORT`, `config.yaml` の `port`, さらに `--port` を一元化 |

```go
func init() {
    f := rootCmd.PersistentFlags()

    // --- CLI フラグ定義 -----------------------------
    // ここで --author / -a フラグなどが CLI に追加される
    f.StringVarP(&cfgFile,    "config",     "c", "",  "Config file path")
    f.StringVarP(&author,     "author",     "a", "",  "Author name")
    f.StringVarP(&projectBase,"projectbase","b", "",  "Base import path")
    f.StringVarP(&userLicense,"license",    "l", "",  "License name")
    f.Bool("viper", true, "Enable Viper")

    // --- Viper にバインド ---------------------------
    // これで Viper のキー author などと CLI フラグがリンク
    // 以降 viper.GetString("author") などで「フラグ値 → なければ他ソース」を自動的に解決してくれる
    viper.BindPFlag("author",      f.Lookup("author"))
    viper.BindPFlag("projectbase", f.Lookup("projectbase"))
    viper.BindPFlag("useViper",    f.Lookup("viper"))

    // --- デフォルト --------------------------------
    // フラグ・ENV・設定ファイルのどれにも値が無ければこのデフォルトが返る
    viper.SetDefault("author",  "NAME HERE <EMAIL>")
    viper.SetDefault("license", "apache")
}

func main() {
    // ...

    // 呼び出し側は以下だけで済む
    name := viper.GetString("author")  // フラグ > ENV > config.yaml > デフォルト
}
```

# cobra使い方 その2: StringVar？StringVarP？String？StringP？

| 関数           | 変数に **格納**するか？         | **ショートハンド** (`-s`) | 戻り値       | 典型的な使い方                |
| ------------ | ---------------------- | ------------------ | --------- | ---------------------- |
| `StringVar`  | **する** (`*string` を渡す) | なし                 | `void`    | 既存の変数に直接セットしたいとき       |
| `StringVarP` | **する** (`*string` を渡す) | **あり**             | `void`    | `-c` など 1 文字の省略形も欲しいとき |
| `String`     | **しない**（内部で変数を作る）      | なし                 | `*string` | 「戻り値でいいから手軽に受け取りたい」とき  |
| `StringP`    | **しない**（内部で変数を作る）      | **あり**             | `*string` | 戻り値 + ショートハンドが欲しいとき    |

* ```go
  StringVarP(&cfgFile, "config", "c", "", "config file path")
  ```

  既存変数 cfgFile に値を書き込み、かつ -c のショートハンドも定義する という意味になる

* Pはpointerのpではない
  * P は Pointer ではなく “shorthand (one-letter) flag を受け取る Variant” を表す記号
  * 由来としては P = “POSIX-style short option”（-c など 1 文字の省略形）と言われている
  * spf13/pflag が Go 標準 flag に “短いフラグ” を追加した際、「短い (= POSIX) 版」を区別する接尾辞 として P を採用

## cobra使い方 その3: 全体の流れ

```go
// File: main.go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir" // ← ~ 展開用（なくても OK）
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/* ----------------------------------------------------------------
   1. アプリ全体で使う設定構造体
-----------------------------------------------------------------*/
type Config struct {
	Author      string `mapstructure:"author"`
	ProjectBase string `mapstructure:"projectbase"`
	License     string `mapstructure:"license"`
	UseViper    bool   `mapstructure:"useViper"`
}

var (
	cfgFile string   // --config の値を格納
	cfg     Config   // viper.Unmarshal で詰め込む構造体
	rootCmd = &cobra.Command{
		Use:   "mytool",
		Short: "Cobra + Viper sample",
		// Cobra は Execute() 時に PersistentPreRunE → RunE の順で呼ぶ。
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig() // 設定ロード & Unmarshal
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			/* ← ここから “メインロジック” に入る想定
			     cfg に値が入った状態で自由に使える */
			fmt.Printf("Loaded config: %+v\n", cfg)
			return nil
		},
	}
)

/* ----------------------------------------------------------------
   2. フラグ定義 & バインド
-----------------------------------------------------------------*/
func init() {
	f := rootCmd.PersistentFlags()

	// --- CLI フラグ --------------------------------------------
	f.StringVarP(&cfgFile,    "config",     "c", "",  "Config file path")
	f.String("author",        "",  "Author name")
	f.StringP("projectbase",  "b", "",  "Base import path (e.g. github.com/you/)")
	f.StringP("license",      "l", "",  "License name")
	f.Bool  ("useViper",      true, "Enable Viper")

	// --- Viper へブリッジ --------------------------------------
	viper.BindPFlag("author",      f.Lookup("author"))
	viper.BindPFlag("projectbase", f.Lookup("projectbase"))
	viper.BindPFlag("useViper",    f.Lookup("useViper"))

	// --- デフォルト値 ------------------------------------------
	viper.SetDefault("author",  "NAME HERE <EMAIL>")
	viper.SetDefault("license", "apache")

	// --- ENV も取り込む（AUTHOR=foo など） ----------------------
	viper.SetEnvPrefix("MYTOOL")                  // AUTHOR → MYTOOL_AUTHOR
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(nil) // ドット→アンダースコア変換不要なら nil
}

/* ----------------------------------------------------------------
   3. 設定ファイル読み込み & 構造体へマッピング
-----------------------------------------------------------------*/
func initConfig() error {
	// A. ファイルパスを決める
	if cfgFile != "" {
		// --config 指定があれば優先
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return fmt.Errorf("resolve home: %w", err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra") // .cobra.yaml 等
	}

	// B. 実際に読む
	if err := viper.ReadInConfig(); err != nil {
		// “設定ファイルがない” 場合は無視しても良い
		if _, notFound := err.(viper.ConfigFileNotFoundError); !notFound {
			return fmt.Errorf("read config: %w", err)
		}
	}

	// C. ファイル/ENV/CLI の最終値を構造体へ
	// StringVar等でバインドしたviperは、configファイルの内容を変数に代入してくれるのか？
	// → 自動では入らない
    //    StringVar で渡した変数に値がセットされるのは pflag がコマンドラインをパースした瞬間だけ 
	//    BindPFlag は「viper 側の世界へ橋渡しする」機能であり、「viper側からの逆流」はしない
	// → 自分で読む
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	// ここで cfg 構造体が完成
	return nil
}

/* ----------------------------------------------------------------
   4. エントリポイント
-----------------------------------------------------------------*/
func main() {
	if err := rootCmd.Execute(); err != nil {
		// Cobra はエラーを返すだけなので明示的に終了させる
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

/* ----------------------------------------------------------------
   5. 補足: 置く YAML ひな形
-------------------------------------------------------------------

# ~/.cobra.yaml など
author: "Taro Yamada <taro@example.com>"
projectbase: "github.com/taro/"
license: "mit"
useViper: true
-----------------------------------------------------------------*/
```

| フィールド  | シグネチャ                                           | 返り値         | エラー伝搬                                                    | 代表的な用途                        |
| ------ | ----------------------------------------------- | ----------- | -------------------------------------------------------- | ----------------------------- |
| `Run`  | `func(cmd *cobra.Command, args []string)`       | **なし**      | 呼び出し元に返せないので、自前で `os.Exit(…)` や `cmd.PrintErrln()` などを行う | “絶対に失敗しない or 失敗しても自分で処理する”ワーク |
| `RunE` | `func(cmd *cobra.Command, args []string) error` | **`error`** | `rootCmd.Execute()` まで伝搬し、テストやラッパから拾える                   | 失敗時に上位へエラーを返したい処理             |

* RunE が設定されていれば優先され、Run は無視される
* 逆に RunE が nil のときだけ Run が呼ばれる
* つまり両方定義しても RunE しか実行されない
  * 可読性のためにも どちらか一方だけ を設定するのが推奨
