package repository

import (
	"example.com/golang-study/domain/model/users_model"
)

type UsersRepository interface {
	FindUser(id string) (*users_model.User, error)
}
