package users

import (
	"example.com/golang-study/model/users_model"
)

func FindUser(id string) (*users_model.User, error) {
	return users_model.FindUser(id)
}
