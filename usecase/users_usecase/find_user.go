package users_usecase

import (
	"example.com/golang-study/domain/model/users_model"
	"example.com/golang-study/infra/users_repository"
)

func FindUser(id string) (*users_model.User, error) {
	repo := users_repository.NewUsersRepository()
	return repo.FindUser(id)
}
