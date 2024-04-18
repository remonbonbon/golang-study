package configutil

import (
	"fmt"
	"io"
	"io/fs"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

type Loader[C any] struct {
	fs fs.FS
}

// Loaderを作成する
func NewLoader[C any](fs_ fs.FS) *Loader[C] {
	return &Loader[C]{fs: fs_}
}

// 設定ファイルを読み込む
func (l *Loader[C]) Load(env Env) (*C, error) {
	conf := map[string]interface{}{}

	// 全環境共通の設定ファイルを読み込む
	{
		c, err := l.loadYAML("config.yaml")
		if err != nil {
			return nil, err
		}
		mergo.Merge(&conf, &c, mergo.WithOverride)
	}

	// 環境用の設定ファイルを読み込む
	{
		c, err := l.loadYAML(fmt.Sprintf("config.%s.yaml", env))
		if err != nil {
			return nil, err
		}
		mergo.Merge(&conf, &c, mergo.WithOverride)
	}

	// mapをC型に変換
	result := new(C)
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

// 設定ファイル(YAML)を読み込む
func (l *Loader[C]) loadYAML(path string) (map[string]interface{}, error) {
	// YAMLファイルを全て読み込む
	f, err := l.fs.Open(path)
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
