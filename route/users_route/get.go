package users_route

import (
	"encoding/json"
	"net/http"
	"regexp"

	res "example.com/golang-study/route/response"
	"example.com/golang-study/usecase/users_usecase"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		res.WriteError(w, r, res.BadRequest("IDは数字です。", nil))
		return
	}

	user, err := users_usecase.FindUser(id)
	if err != nil {
		switch err.(type) {
		// case *users_model.NotFoundError:
		// 	res.WriteError(w, r, res.NotFound("ユーザーが見つかりません。", nil))
		default:
			res.WriteError(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(user)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}
	w.Write(bytes)
}
