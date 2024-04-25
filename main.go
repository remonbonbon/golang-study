package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/golang-study/common"
	"example.com/golang-study/config"
	"example.com/golang-study/web"
)

func main() {
	conf := config.Get()
	slog.SetDefault(common.NewLogger())
	srv := http.Server{Addr: conf.Listen, Handler: web.NewRouter()}

	slog.Debug("でばっぐ", slog.String("hoge", "piyo"))
	slog.Info("いんふぉ", slog.String("hoge", "piyo"))
	slog.Warn("わーん", slog.String("hoge", "piyo"))
	slog.Error("えらー", slog.String("hoge", "piyo"))
	slog.Info("グループテスト", slog.Group("g", slog.Int("a", 1), slog.Int("b", 2)))
	slog.Info("グループ入れ子", slog.Group("A", slog.Group("B", slog.Int("a", 1)), slog.Int("b", 2)))
	slog.With(slog.Group("group", slog.Int("hoge", 123))).Info("Withテスト")

	// Graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		// SIGINT or SIGTERMを待機
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		// シグナル発生したらサーバーを終了
		if err := srv.Shutdown(context.Background()); err != nil {
			slog.Error("Shutdown failed", slog.Any("error", err))
		}
		slog.Info("Shutdown server")
		close(idleConnsClosed)
	}()

	// サーバー起動
	slog.Info(fmt.Sprintf("Listen on http://%s", conf.Listen))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("Listen failed", slog.Any("error", err))
		return
	}

	// Graceful shutdown完了を待機
	<-idleConnsClosed
}
