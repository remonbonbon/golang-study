package web

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"example.com/golang-study/common"
	res "example.com/golang-study/web/response"
	"example.com/golang-study/web/users"
	"example.com/golang-study/web/welcome"
)

// ルーティングを設定
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// ミドルウェア
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(LoggerMiddleware)
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
	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		res.ErrorJson(w, r, common.SystemError("エラーテスト", errors.New("test error")))
	})

	return r
}
