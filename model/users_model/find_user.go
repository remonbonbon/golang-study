package users_model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FindUser(id string) (*User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, &HTTPError{Req: req}
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, &HTTPError{Req: req, Res: resp, Err: err}
	}

	if resp.StatusCode == 404 {
		return nil, &NotFoundError{}
	}
	if resp.StatusCode != 200 {
		return nil, &HTTPError{Req: req, Res: resp}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &HTTPError{Req: req, Res: resp, Err: err}
	}

	var u User
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, &HTTPError{Req: req, Res: resp, Err: err}
	}
	return &u, nil
}
