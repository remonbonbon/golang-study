package configutil

import (
	"flag"
	"os"
)

type Env string

const (
	EnvTest Env = "test"
	EnvDev  Env = "dev"
	EnvStg  Env = "stg"
	EnvProd Env = "prod"
)

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
