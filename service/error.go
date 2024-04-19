package service

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

type BusinessError struct {
	Status  int    `json:"status"`  // HTTPステータスコード
	Message string `json:"message"` // エラーメッセージ
	Error   *error `json:"error"`   // 発生したエラー (任意)
}

// HTTPステータスコード 400 Bad Request
func BadRequest(m string, e *error) BusinessError {
	return BusinessError{Status: 400, Message: m, Error: e}
}

// HTTPステータスコード 404 Not Found
func NotFound(m string, e *error) BusinessError {
	return BusinessError{Status: 404, Message: m, Error: e}
}

// HTTPステータスコード 500 Internal Server Error
func InternalServerError(m string, e *error) BusinessError {
	return BusinessError{Status: 500, Message: m, Error: e}
}

type errorResponse struct {
	Message string `json:"message"` // エラーメッセージ
}

// エラーレスポンスを送信する
func WriteError(w http.ResponseWriter, r *http.Request, e BusinessError) {
	logger := httplog.LogEntry(r.Context())

	// ステータスコード未設定の場合、500にする
	if e.Status < 100 {
		e.Status = 500
	}

	// ステータスコード400番台はログレベル WARN、500番台はログレベル ERROR
	if e.Status < 500 {
		logger.Warn(e.Message, slog.Any("error", e.Error))
	} else {
		logger.Error(e.Message, slog.Any("error", e.Error))
	}

	j, err := json.Marshal(errorResponse{Message: e.Message})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	w.Write([]byte(j))
}
