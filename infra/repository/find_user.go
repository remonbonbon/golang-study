package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/golang-study/domain/model"
	"example.com/golang-study/infra"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (repo *UsersRepository) FindUser(id string) (*model.User, bool, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, false, infra.NewHTTPError("request failed", req, nil, err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, false, infra.NewHTTPError("request failed", req, resp, err)
	}

	if resp.StatusCode == 404 {
		return nil, false, nil
	}
	if resp.StatusCode != 200 {
		return nil, false, infra.NewHTTPError("unexpected response", req, resp, nil)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, infra.NewHTTPError("read body failed", req, resp, err)
	}

	var u model.User
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, false, infra.NewHTTPError("json failed", req, resp, err)
	}
	return &u, true, nil
}
