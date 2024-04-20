package infra

import (
	"fmt"
	"net/http"

	"example.com/golang-study/common"
)

type HTTPError struct {
	Req *http.Request
	Res *http.Response
	Err error
}

func (e *HTTPError) Error() string {
	var str string
	str += fmt.Sprintf("req=%+v", e.Req)
	str += ", "
	str += fmt.Sprintf("res=%+v", e.Res)
	str += ", "
	str += fmt.Sprintf("err=%+v", e.Err)

	return str
}

// ユーザーの入力がおかしい場合のエラー
func SystemError(m string, req *http.Request, res *http.Response, e error) *common.BusinessError {
	return common.SystemError(m, &HTTPError{Req: req, Err: e})
}

// 指定されたものが見つからなかった場合のエラー
func NotFound(m string, req *http.Request, res *http.Response, e error) *common.BusinessError {
	return common.NotFound(m, &HTTPError{Req: req, Err: e})
}
