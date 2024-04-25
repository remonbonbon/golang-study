package web

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"example.com/golang-study/common"
	"example.com/golang-study/web/users"
	"example.com/golang-study/web/welcome"
)

func MyLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := common.LogWith(ctx)

		log.Info("Before")
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Info("After")
	}
	return http.HandlerFunc(fn)
}

// ルーティングを設定
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// ミドルウェア
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(MyLogger)
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
