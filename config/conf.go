package config

// 設定値
// - 全てpublic(先頭を大文字)にする必要がある。
// - 変数定義の最後に `yaml:"fuga"` を付けるとキーを変更できる。
// See: https://pkg.go.dev/gopkg.in/yaml.v3#Marshal
type Config struct {
	Listen   string // listenするホストとポート (例: 0.0.0.0:8080)
	LogLevel string // ログレベル (DEBUG, INFO, WARN, ERROR)

	Fuga struct {
		Bar struct {
			Value1 int
			Value2 int
		}
	}
}
