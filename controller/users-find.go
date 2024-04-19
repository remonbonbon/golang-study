package controller

import (
	"net/http"
	"regexp"

	"example.com/golang-study/model/dummy"
)

func FindUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		WriteError(w, r, BadRequest("IDは数字です。", nil))
		return
	}

	body, err := dummy.FindUser(id)
	if err != nil {
		WriteError(w, r, InternalServerError("エラーです。", &err))
		return
	}

	s := string(body)
	if s == "{}" {
		WriteError(w, r, NotFound("見つかりませんでした。", &err))
		return
	}

	w.Write([]byte(body))
}
