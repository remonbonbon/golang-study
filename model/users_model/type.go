package users_model

import (
	"fmt"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

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

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "not found"
}
