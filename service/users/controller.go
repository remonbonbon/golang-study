package users

import (
	"net/http"
	"regexp"

	srv "example.com/golang-study/service"
	"example.com/golang-study/service/users/dummymodel"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		srv.WriteError(w, r, srv.BadRequest("IDは数字です。", nil))
		return
	}

	body, err := dummymodel.FindUser(id)
	if err != nil {
		srv.WriteError(w, r, srv.InternalServerError("エラーです。", &err))
		return
	}

	s := string(body)
	if s == "{}" {
		srv.WriteError(w, r, srv.NotFound("見つかりませんでした。", &err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
