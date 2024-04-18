package configutil

import (
	"io/fs"
	"sync"
)

// プログラム全体からアクセスする設定値
type GlobalLoader[C any] struct {
	loader *Loader[C] // 実際のLoader
	conf   *C         // 設定値
	mutex  sync.Mutex
}

// GlobalLoaderを作成する
func NewGlobalLoader[C any](fs_ fs.FS) *GlobalLoader[C] {
	return &GlobalLoader[C]{
		loader: NewLoader[C](fs_),
	}
}

// 設定ファイルを読み込む
func (g *GlobalLoader[C]) Load() error {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	env := CurrentEnv()
	conf, err := g.loader.Load(env)
	if err != nil {
		return err
	}
	g.conf = conf

	return nil
}

// 設定値を返す。
// 設定ファイルが読み込まれていない場合は読み込む。
func (g *GlobalLoader[C]) Get() *C {
	if g.conf == nil {
		err := g.Load()
		if err != nil {
			panic(err)
		}
	}
	return g.conf
}
