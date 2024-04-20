package repository

import (
	"fmt"

	"example.com/golang-study/domain/model"
	"example.com/golang-study/infra/internal/http"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repo *UsersRepository) FindUser(id string) (*model.User, bool, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	return http.HTTPGet[model.User](url)
}
