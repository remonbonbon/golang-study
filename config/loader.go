package config

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

// 設定ファイルをバイナリに埋め込む。
// 一番下のコメントディレクティブで指定したファイルが埋め込まれる。
//
//go:embed *.yaml
var configFS embed.FS

// 設定ファイルを読み込む。
//
// 読み込み順序:
// 1. config.yaml
// 2. config.local.yaml
// 3. config.${ENV}.yaml
// 4. config.${ENV}.local.yaml
func LoadConfig(env Env) (*Config, error) {
	conf := map[string]interface{}{}

	// 読み込み順に設定ファイルを読み込んで上書きマージしていく
	for _, path := range []string{
		"config.yaml",
		"config.local.yaml",
		fmt.Sprintf("config.%s.yaml", env),
		fmt.Sprintf("config.%s.local.yaml", env),
	} {
		c, err := loadYAML(configFS, path)
		// ～.local.yamlは存在しなくてもOK。それ以外（config.yamlが無い場合等）はエラーにする
		if strings.Contains(path, ".local.") && os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		mergo.Merge(&conf, &c, mergo.WithOverride)
	}

	// mapをC型に変換
	result := new(Config)
	tmp, err := yaml.Marshal(conf)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(tmp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// YAMLファイルを読み込む
func loadYAML(fs_ fs.FS, path string) (map[string]interface{}, error) {
	// YAMLファイルを全て読み込む
	f, err := fs_.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// YAMLをパース
	conf := map[string]interface{}{}
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
