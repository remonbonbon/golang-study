package common

import (
	"net/http"
	"runtime"
)

type BusinessError struct {
	Status  int    `json:"status"`  // HTTPステータスコード
	Message string `json:"message"` // エラーメッセージ
	Err     error  `json:"error"`   // 発生したエラー (任意)
	File    string `json:"file"`    // 呼び出し元のファイル
	Line    int    `json:"line"`    // 呼び出し元の行
}

func (e *BusinessError) Error() string {
	return e.Message
}

// ユーザーの入力がおかしい場合のエラー。
// HTTPステータスコード 400 Bad Request
func InvalidInput(m string, e error) *BusinessError {
	b := BusinessError{Status: http.StatusBadRequest, Message: m, Err: e}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		b.File = file
		b.Line = line
	}
	return &b
}

// 指定されたものが見つからなかった場合のエラー。
// HTTPステータスコード 404 Not Found
func NotFound(m string, e error) *BusinessError {
	b := BusinessError{Status: http.StatusNotFound, Message: m, Err: e}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		b.File = file
		b.Line = line
	}
	return &b
}

// システムエラー。
// HTTPステータスコード 500 Internal Server Error
func SystemError(m string, e error) *BusinessError {
	b := BusinessError{Status: http.StatusInternalServerError, Message: m, Err: e}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		b.File = file
		b.Line = line
	}
	return &b
}
