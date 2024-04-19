package config

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

// 環境
type Env string

const (
	EnvTest Env = "test" // go test実行時の環境
	EnvDev  Env = "dev"  // dev環境
	EnvStg  Env = "stg"  // staging環境
	EnvProd Env = "prod" // production環境
)

// プログラム全体からアクセスする設定値
type globalConfig struct {
	conf  *Config
	mutex sync.Mutex
}

var g globalConfig

// 設定値を取得
func Get() *Config {
	if g.conf == nil {
		Load()
	}
	return g.conf
}

// 設定ファイルを読み込む。
//
// 読み込み順序:
// 1. config.yaml
// 2. config.local.yaml
// 3. config.${ENV}.yaml
// 4. config.${ENV}.local.yaml
func Load() {
	// 設定ファイル読み込み
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	env := CurrentEnv()
	defaultConfig("config")
	localConfig("config.local")
	envConfig(fmt.Sprintf("config.%s", env))
	localConfig(fmt.Sprintf("config.%s.local", env))
	// fmt.Printf("%+v\n", viper.AllSettings())

	// パースしてConfigに詰め込む
	defer g.mutex.Unlock()
	g.mutex.Lock()

	g.conf = new(Config)
	if err := viper.Unmarshal(g.conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	// fmt.Printf("%+v\n", *g.conf)
}

// デフォルト設定ファイルを読み込む。無い場合はエラー
func defaultConfig(name string) {
	viper.SetConfigName(name)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// ローカル用設定ファイルをマージする。無くてもOK
func localConfig(name string) {
	viper.SetConfigName(name)

	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}

// 各環境用設定ファイルをマージする。無い場合はエラー
func envConfig(name string) {
	viper.SetConfigName(name)

	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// 現在の環境を返す
func CurrentEnv() Env {
	if IsTestEnv() {
		return EnvTest
	}
	env := os.Getenv("ENV")
	if env == "" {
		// ローカル環境もdev環境扱いとする
		return EnvDev
	}
	return Env(env)
}

// テスト環境の場合trueを返す
func IsTestEnv() bool {
	// go testでTestMain()が実行されるタイミングで設定される
	return flag.Lookup("test.v") != nil
}

// dev環境の場合trueを返す
func IsDevEnv() bool {
	return CurrentEnv() == EnvDev
}
