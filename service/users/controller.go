package users

import (
	"encoding/json"
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

	user, err := dummymodel.FindUser(id)
	if err != nil {
		switch err.(type) {
		case *dummymodel.NotFoundError:
			srv.WriteError(w, r, srv.NotFound("ユーザーが見つかりません。", nil))
		default:
			srv.WriteError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(user)
	if err != nil {
		srv.WriteError(w, r, err)
		return
	}
	w.Write(bytes)
}
