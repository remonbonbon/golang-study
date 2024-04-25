package users

import (
	"net/http"
	"regexp"

	"example.com/golang-study/common"
	"example.com/golang-study/infra/repository"
	"example.com/golang-study/usecase/users"
	res "example.com/golang-study/web/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	if !re.MatchString(id) {
		res.ErrorJson(w, r, common.InvalidInput("IDは数字です。", nil))
		return
	}

	repo := repository.NewUsersRepository()
	user, err := users.FindUser(repo, id)
	if err != nil {
		res.ErrorJson(w, r, err)
		return
	}

	res.Json(w, r, user)
}
