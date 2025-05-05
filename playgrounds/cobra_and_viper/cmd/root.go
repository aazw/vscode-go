package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// 1. アプリ全体で使う設定構造体
type Config struct {
	Author      string `mapstructure:"author"`
	ProjectBase string `mapstructure:"projectbase"`
	License     string `mapstructure:"license"`
	UseViper    bool   `mapstructure:"useViper"`
}

var (
	cfgFile string // --config の値を格納
	cfg     Config // viper.Unmarshal で詰め込む構造体
	rootCmd = &cobra.Command{
		Use:   "mytool",
		Short: "Cobra + Viper sample",
		// Cobra は Execute() 時に PersistentPreRunE → RunE の順で呼ぶ。
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig() // 設定ロード & Unmarshal
		},
		RunE: runE,
	}
)

// 2. フラグ定義 & バインド
func init() {
	f := rootCmd.PersistentFlags()

	// --- CLI フラグ --------------------------------------------
	f.StringVarP(&cfgFile, "config", "c", "", "Config file path")
	f.String("author", "", "Author name")
	f.StringP("projectbase", "b", "", "Base import path (e.g. github.com/you/)")
	f.StringP("license", "l", "", "License name")
	f.Bool("useViper", true, "Enable Viper")

	// --- Viper へブリッジ --------------------------------------
	viper.BindPFlag("author", f.Lookup("author"))
	viper.BindPFlag("projectbase", f.Lookup("projectbase"))
	viper.BindPFlag("useViper", f.Lookup("useViper"))

	// --- デフォルト値 ------------------------------------------
	viper.SetDefault("author", "NAME HERE <EMAIL>")
	viper.SetDefault("license", "apache")

	// --- ENV も取り込む（AUTHOR=foo など） ----------------------
	viper.SetEnvPrefix("MYTOOL") // AUTHOR → MYTOOL_AUTHOR
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// 3. 設定ファイル読み込み & 構造体へマッピング
func initConfig() error {
	// A. ファイルパスを決める
	if cfgFile != "" {
		// --config 指定があれば優先
		viper.SetConfigFile(cfgFile)
	} else {
		// // use home directory
		// homeDir, err := os.UserHomeDir()
		// if err != nil {
		// 	return fmt.Errorf("resolve home: %w", err)
		// }
		// viper.AddConfigPath(homeDir)

		// use current directory
		cd, _ := os.Getwd()
		viper.AddConfigPath(cd)

		viper.SetConfigName(".cobra") // .cobra.yaml 等
	}

	// B. 実際に読む
	if err := viper.ReadInConfig(); err != nil {
		// “設定ファイルがない” 場合は無視しても良い
		if _, notFound := err.(*viper.ConfigFileNotFoundError); notFound {
			return fmt.Errorf("read config: %w", err)
		} else {
			return fmt.Errorf("unknown error: %w", err)
		}
	}

	// C. ファイル/ENV/CLI の最終値を構造体へ
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	// ここで cfg 構造体が完成
	return nil
}

// 4. メインロジック. cfg に値が入った状態で自由に使える
func runE(cmd *cobra.Command, args []string) error {

	fmt.Printf("Loaded config: %+v\n", cfg)

	fmt.Println()

	buf, _ := yaml.Marshal(cfg)
	fmt.Printf("%s\n", string(buf))

	return nil
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("rootCmd: %w", err)
	}
	return nil
}
