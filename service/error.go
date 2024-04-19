package service

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

type ServiceError struct {
	Status  int    // HTTPステータスコード
	Message string // エラーメッセージ
	Error   *error // 発生したエラー (任意)
}

// HTTPステータスコード 400 Bad Request
func BadRequest(m string, e *error) ServiceError {
	return ServiceError{Status: 400, Message: m, Error: e}
}

// HTTPステータスコード 404 Not Found
func NotFound(m string, e *error) ServiceError {
	return ServiceError{Status: 404, Message: m, Error: e}
}

// HTTPステータスコード 500 Internal Server Error
func InternalServerError(m string, e *error) ServiceError {
	return ServiceError{Status: 500, Message: m, Error: e}
}

// エラーレスポンスを送信する
func WriteError(w http.ResponseWriter, r *http.Request, e ServiceError) {
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

	j, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(e.Status)
	w.Write([]byte(j))
}
