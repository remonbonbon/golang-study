package repository

import (
	"example.com/golang-study/domain/model"
)

type UsersRepository interface {
	FindUser(id string) (*model.User, bool, error)
}
