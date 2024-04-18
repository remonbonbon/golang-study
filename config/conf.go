package config

import (
	"embed"

	"example.com/golang-study/config/configutil"
)

// 設定値
// - 全てpublic(先頭を大文字)にする必要がある
// - 変数定義の最後に `yaml:"fuga"` を付けるとキーを変更できる
// See: https://pkg.go.dev/gopkg.in/yaml.v3#Marshal
type Config struct {
	Hoge string
	Fuga struct {
		Bar struct {
			Value1 int
			Value2 int
		}
	}
}

// 設定ファイルをバイナリに埋め込む
// 一番下のコメントディレクティブで指定したファイルが埋め込まれる
//
//go:embed *.yaml
var ConfigFS embed.FS

// 設定値へグローバルにアクセスできるようにする
var globalLoader = configutil.NewGlobalLoader[Config](ConfigFS)

func Load() error  { return globalLoader.Load() }
func Get() *Config { return globalLoader.Get() }
