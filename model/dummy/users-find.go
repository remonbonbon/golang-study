package dummy

import (
	"fmt"
	"io"
	"net/http"
)

func FindUser(id string) ([]byte, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
