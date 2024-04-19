package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"example.com/golang-study/config"
	"example.com/golang-study/server"
)

func main() {
	conf := config.Get()

	logger := server.NewLogger()
	r := server.NewRouter(logger)

	// サーバー起動
	logger.Info(fmt.Sprintf("Listen on http://%s", conf.Listen))
	err := http.ListenAndServe(conf.Listen, r)
	if err != nil {
		logger.Error("Listen failed", slog.Any("error", err))
	}
}
