package route

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"example.com/golang-study/common"
	"github.com/go-chi/httplog/v2"
)

type errorResponse struct {
	Message string `json:"message"` // エラーメッセージ
}

// エラーレスポンスを送信する
func WriteError(w http.ResponseWriter, r *http.Request, originalError error) {
	logger := httplog.LogEntry(r.Context())

	var e *common.BusinessError
	switch be := originalError.(type) {
	case *common.BusinessError:
		e = be
	default:
		// BusinessError以外の場合
		e = &common.BusinessError{Status: 500, Message: "system error", Err: originalError}
	}

	// ステータスコード未設定の場合、500にする
	if e.Status < 100 {
		e.Status = 500
	}

	// ステータスコード400番台はログレベル WARN、500番台はログレベル ERROR
	var attrs []any
	if e.Err != nil {
		attrs = append(attrs, slog.Any("error", e.Err))
	}
	if e.Status < 500 {
		logger.Warn(e.Message, attrs...)
	} else {
		logger.Error(e.Message, attrs...)
	}

	j, err := json.Marshal(errorResponse{Message: e.Message})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	w.Write([]byte(j))
}
