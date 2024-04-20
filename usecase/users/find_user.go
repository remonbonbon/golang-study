package users

import (
	"example.com/golang-study/domain/model"
	"example.com/golang-study/infra/repository"
)

func FindUser(id string) (*model.User, error) {
	repo := repository.NewUsersRepository()
	return repo.FindUser(id)
}
