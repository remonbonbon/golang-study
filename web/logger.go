package web

import (
	"fmt"
	"net/http"

	"example.com/golang-study/common"
)

// ロガーをcontextに埋め込む、アクセスログを出力する
func LoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := common.ContextWithLogger(r.Context()) // ロガーをcontextに埋め込む
		log := common.LogWith(ctx)

		method := r.Method
		path := r.URL.Path

		// リクエストのログ
		log.Info(fmt.Sprintf("%s %s", method, path))

		// リクエストの処理
		sw := &statusWriter{OriginalWriter: w}
		next.ServeHTTP(sw, r.WithContext(ctx))

		// レスポンスのログ
		log.Info(fmt.Sprintf("%s %s (%d)", method, path, sw.Status))
	}
	return http.HandlerFunc(fn)
}

// 書き込んだステータスコードを保存するResponseWriter
type statusWriter struct {
	OriginalWriter http.ResponseWriter
	Status         int
}

func (w *statusWriter) Header() http.Header {
	return w.OriginalWriter.Header()
}
func (w *statusWriter) Write(b []byte) (int, error) {
	return w.OriginalWriter.Write(b)
}
func (w *statusWriter) WriteHeader(status int) {
	w.Status = status
	w.OriginalWriter.WriteHeader(status)
}
