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
	"example.com/golang-study/controller"
)

func main() {
	// 設定ファイル読み込み
	config.Load()
	conf := config.Get()

	// Logger
	var logger *httplog.Logger
	{
		devMode := config.IsDevEnv()
		logLevel := map[string]slog.Level{
			"DEBUG": slog.LevelDebug,
			"INFO":  slog.LevelInfo,
			"WARN":  slog.LevelWarn,
			"ERROR": slog.LevelError,
		}
		source := "source"
		if devMode {
			source = "" // Set "" to disable
		}
		logger = httplog.NewLogger("main", httplog.Options{
			LogLevel:         logLevel[conf.LogLevel],
			MessageFieldName: "message",
			Concise:          devMode,
			JSON:             !devMode,
			SourceFieldName:  source,
		})
	}

	r := chi.NewRouter()

	// ミドルウェア
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)

	// ルーティング
	r.Get("/", controller.Welcome)
	r.Get("/users/{id}", controller.FindUser)

	// サーバー起動
	logger.Debug("config", slog.Any("config", *conf))
	logger.Info(fmt.Sprintf("Listen on http://%s", conf.Listen))
	err := http.ListenAndServe(conf.Listen, r)
	if err != nil {
		logger.Error("Listen failed", slog.Any("error", err))
	}
}
