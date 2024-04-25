package response

import (
	"log/slog"
	"net/http"

	"example.com/golang-study/common"
	"example.com/golang-study/common/message"
)

type ErrorResponse struct {
	Message string `json:"message"` // エラーメッセージ
}

// エラーレスポンスを送信する
func ErrorJson(w http.ResponseWriter, r *http.Request, originalError error) {
	log := common.LogWith(r.Context())

	// BusinessErrorの場合はそのステータスコード等を使用する
	var e *common.BusinessError
	switch be := originalError.(type) {
	case *common.BusinessError:
		e = be
	default:
		// BusinessError以外の場合
		e = &common.BusinessError{Status: http.StatusInternalServerError, Message: message.SystemError(), Err: originalError}
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
		log.Warn(e.Message, attrs...)
	} else {
		log.Error(e.Message, attrs...)
	}

	JsonWithStatus(w, r, ErrorResponse{Message: e.Message}, e.Status)
}
