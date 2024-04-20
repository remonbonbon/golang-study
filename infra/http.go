package infra

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Message  string
	Request  *http.Request
	Response *http.Response
	Err      error
}

func (e *HTTPError) Error() string {
	var str string
	str += e.Message
	str += ", "
	str += fmt.Sprintf("req=%+v", e.Request)
	str += ", "
	str += fmt.Sprintf("res=%+v", e.Response)
	str += ", "
	str += fmt.Sprintf("err=%+v", e.Err)

	return str
}

func NewHTTPError(m string, req *http.Request, res *http.Response, e error) *HTTPError {
	return &HTTPError{Message: m, Request: req, Response: res, Err: e}
}
