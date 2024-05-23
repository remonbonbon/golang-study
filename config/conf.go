package config

// 設定値
// - 全てpublic(先頭を大文字)にする必要がある。
// - 変数定義の最後に `yaml:"fuga"` を付ける
// See: https://pkg.go.dev/gopkg.in/yaml.v3#Marshal
type Config struct {
	Listen   string `yaml:"listen"`   // listenするホストとポート (例: 0.0.0.0:8080)
	LogLevel string `yaml:"logLevel"` // ログレベル (DEBUG, INFO, WARN, ERROR)

	Hoge string `yaml:"hoge"`
	Fuga struct {
		Bar struct {
			Value1 int `yaml:"value1"`
			Value2 int `yaml:"value2"`
		} `yaml:"bar"`
	} `yaml:"fuga"`
}
