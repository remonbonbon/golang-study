package users_controller

import (
	"encoding/json"
	"net/http"
	"regexp"

	ctl "example.com/golang-study/controller"
	"example.com/golang-study/model/users_model"
	"example.com/golang-study/service/users"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		ctl.WriteError(w, r, ctl.BadRequest("IDは数字です。", nil))
		return
	}

	user, err := users.FindUser(id)
	if err != nil {
		switch err.(type) {
		case *users_model.NotFoundError:
			ctl.WriteError(w, r, ctl.NotFound("ユーザーが見つかりません。", nil))
		default:
			ctl.WriteError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(user)
	if err != nil {
		ctl.WriteError(w, r, err)
		return
	}
	w.Write(bytes)
}
