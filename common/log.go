package common

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5/middleware"

	"example.com/golang-study/config"
)

// ロガーを作成
func NewLogger() *slog.Logger {
	conf := config.Get()
	devMode := config.IsDevEnv()
	logLevel := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}

	opt := &slog.HandlerOptions{
		AddSource: !devMode,
		Level:     logLevel[conf.LogLevel]}

	var logger *slog.Logger
	if devMode {
		// logger = slog.New(slog.NewTextHandler(os.Stdout, opt))
		logger = slog.New(NewMyHandler(os.Stdout, &MyHandlerOptions{Level: slog.LevelDebug}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opt))
	}
	return logger
}

func LogWith(ctx context.Context) *slog.Logger {
	var l = slog.Default()
	l = l.With(
		slog.String("reqID1", middleware.GetReqID(ctx)),
		slog.String("reqID2", middleware.GetReqID(ctx)))
	return l
}
