package users_repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/golang-study/domain/model/users_model"
	"example.com/golang-study/infra"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repo *UsersRepository) FindUser(id string) (*users_model.User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, &infra.HTTPError{Req: req}
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, &infra.HTTPError{Req: req, Res: resp, Err: err}
	}

	if resp.StatusCode == 404 {
		return nil, &infra.NotFoundError{}
	}
	if resp.StatusCode != 200 {
		return nil, &infra.HTTPError{Req: req, Res: resp}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &infra.HTTPError{Req: req, Res: resp, Err: err}
	}

	var u users_model.User
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, &infra.HTTPError{Req: req, Res: resp, Err: err}
	}
	return &u, nil
}
