package server

import (
	"log/slog"

	"github.com/go-chi/httplog/v2"

	"example.com/golang-study/config"
)

// ロガーを作成
func NewLogger() *httplog.Logger {
	conf := config.Get()

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
	return logger
}
