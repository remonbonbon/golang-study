package users_route

import (
	"encoding/json"
	"net/http"
	"regexp"

	"example.com/golang-study/common"
	res "example.com/golang-study/route/response"
	"example.com/golang-study/usecase/users"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		res.WriteError(w, r, common.InvalidInput("IDは数字です。", nil))
		return
	}

	user, err := users.FindUser(id)
	if err != nil {
		res.WriteError(w, r, err)
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
