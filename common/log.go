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

	var logger *slog.Logger
	if devMode {
		// logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// 	AddSource: !devMode,
		// 	Level:     logLevel[conf.LogLevel]}))
		logger = slog.New(NewHumanHandler(os.Stdout, &HumanHandlerOptions{
			Level: slog.LevelDebug}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: !devMode,
			Level:     logLevel[conf.LogLevel]}))
	}
	return logger
}

func LogWith(ctx context.Context) *slog.Logger {
	var l = slog.Default()
	l = l.With(
		slog.String("reqID", middleware.GetReqID(ctx)))
	return l
}
