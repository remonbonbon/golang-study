package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"example.com/golang-study/common"
	"example.com/golang-study/web/users"
	"example.com/golang-study/web/welcome"
)

// ルーティングを設定
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// ミドルウェア
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(myLogger)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)

	// 404カスタムハンドラ
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		// 空のレスポンスボディを返す
	})

	// ルーティング
	r.Get("/", welcome.Index)
	r.Get("/users/{id}", users.Get)

	return r
}

// ロガーをcontextに埋め込む、アクセスログを出力する
func myLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := common.ContextWithLogger(r.Context()) // ロガーをcontextに埋め込む
		log := common.LogWith(ctx)

		method := r.Method
		path := r.URL.Path

		// リクエストのログ
		log.Info(fmt.Sprintf("%s %s", method, path))

		// リクエストの処理
		sw := &StatusWriter{OriginalWriter: w}
		next.ServeHTTP(sw, r.WithContext(ctx))

		// レスポンスのログ
		log.Info(fmt.Sprintf("%s %s (%d)", method, path, sw.Status))
	}
	return http.HandlerFunc(fn)
}

// 書き込んだステータスコードを保存するResponseWriter
type StatusWriter struct {
	OriginalWriter http.ResponseWriter
	Status         int
}

func (w *StatusWriter) Header() http.Header {
	return w.OriginalWriter.Header()
}
func (w *StatusWriter) Write(b []byte) (int, error) {
	return w.OriginalWriter.Write(b)
}
func (w *StatusWriter) WriteHeader(status int) {
	w.Status = status
	w.OriginalWriter.WriteHeader(status)
}
