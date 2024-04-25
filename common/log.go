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

type ctxKey int

const (
	loggerKey ctxKey = iota
)

// Contextにロガーを保存
func ContextWithLogger(ctx context.Context) context.Context {
	log := slog.Default().With(
		slog.String("reqID", middleware.GetReqID(ctx)))
	return context.WithValue(ctx, loggerKey, log)
}

// Contextのロガーを取得
func LogWith(ctx context.Context) *slog.Logger {
	log, ok := ctx.Value(loggerKey).(*slog.Logger)
	if ok {
		return log
	}
	return slog.Default()
}
