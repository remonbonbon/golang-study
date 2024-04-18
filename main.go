package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"

	"example.com/golang-study/config"
)

func main() {
	fmt.Printf("%+v\n", *config.Get())

	r := chi.NewRouter()

	// Logger
	devMode := true
	logger := httplog.NewLogger("main", httplog.Options{
		LogLevel:         slog.LevelDebug,
		MessageFieldName: "message",
		Concise:          devMode,
		JSON:             !devMode,
	})

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Compress(5))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		w.Write([]byte("welcome"))
		oplog.Info("ログてすと")
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("エラーです！")
	})

	http.ListenAndServe("localhost:8080", r)
}
