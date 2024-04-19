package service

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

type BusinessError struct {
	Status  int    `json:"status"`  // HTTPステータスコード
	Message string `json:"message"` // エラーメッセージ
	Err     error  `json:"error"`   // 発生したエラー (任意)
}

func (e *BusinessError) Error() string {
	return e.Message
}

// HTTPステータスコード 400 Bad Request
func BadRequest(m string, e error) *BusinessError {
	return &BusinessError{Status: 400, Message: m, Err: e}
}

// HTTPステータスコード 404 Not Found
func NotFound(m string, e error) *BusinessError {
	return &BusinessError{Status: 404, Message: m, Err: e}
}

// HTTPステータスコード 500 Internal Server Error
func InternalServerError(m string, e error) *BusinessError {
	return &BusinessError{Status: 500, Message: m, Err: e}
}

type errorResponse struct {
	Message string `json:"message"` // エラーメッセージ
}

// エラーレスポンスを送信する
func WriteError(w http.ResponseWriter, r *http.Request, e error) {
	if e == nil {
		panic(errors.New("unexpected nil error"))
	}
	logger := httplog.LogEntry(r.Context())

	var e2 *BusinessError
	switch myErr := e.(type) {
	case *BusinessError:
		e2 = myErr
	default:
		// BusinessError以外の場合
		e2 = &BusinessError{Status: 500, Message: "システムエラー", Err: e}
	}

	// ステータスコード未設定の場合、500にする
	if e2.Status < 100 {
		e2.Status = 500
	}

	// ステータスコード400番台はログレベル WARN、500番台はログレベル ERROR
	var attrs []any
	if e2.Err != nil {
		attrs = append(attrs, slog.Any("error", e2.Err))
	}
	if e2.Status < 500 {
		logger.Warn(e2.Message, attrs...)
	} else {
		logger.Error(e2.Message, attrs...)
	}

	j, err := json.Marshal(errorResponse{Message: e2.Message})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e2.Status)
	w.Write([]byte(j))
}
