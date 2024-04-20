package log

import (
	"log/slog"
	"time"

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
			Concise:          devMode,
			JSON:             !devMode,
			TimeFieldName:    "timestamp",
			TimeFieldFormat:  time.RFC3339,
			MessageFieldName: "message",
			LevelFieldName:   "level",
			SourceFieldName:  source,
		})
	}
	return logger
}
