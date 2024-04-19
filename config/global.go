package config

import (
	"flag"
	"fmt"
	"os"
	"sync"
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
		defer g.mutex.Unlock()
		g.mutex.Lock()

		env := CurrentEnv()
		conf, err := LoadConfig(env)
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		g.conf = conf
	}
	return g.conf
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
