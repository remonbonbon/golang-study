package users

import (
	"net/http"
	"regexp"

	"example.com/golang-study/common"
	"example.com/golang-study/infra/repository"
	"example.com/golang-study/logic/users"
	res "example.com/golang-study/ui/internal/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		res.WriteError(w, r, common.InvalidInput("IDは数字です。", nil))
		return
	}

	repo := repository.NewUsersRepository()
	user, err := users.FindUser(repo, id)
	if err != nil {
		res.WriteError(w, r, err)
		return
	}

	res.Json(w, r, user)
}
