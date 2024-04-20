package ui

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"

	"example.com/golang-study/ui/users"
	"example.com/golang-study/ui/welcome"
)

// ルーティングを設定
func NewRouter(logger *httplog.Logger) *chi.Mux {
	r := chi.NewRouter()

	// ミドルウェア
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)

	// ルーティング
	r.Get("/", welcome.Index)
	r.Get("/users/{id}", users.Get)

	return r
}
