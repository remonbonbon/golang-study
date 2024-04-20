package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func HTTPGet[Body any](url string) (*Body, bool, error) {
	return HTTPRequest[Body]("GET", url)
}

func HTTPRequest[Body any](method string, url string) (*Body, bool, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, false, NewHTTPError("request failed", req, nil, err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, false, NewHTTPError("request failed", req, resp, err)
	}

	if resp.StatusCode == 404 {
		return nil, false, nil
	}
	if resp.StatusCode != 200 {
		return nil, false, NewHTTPError("unexpected response", req, resp, nil)
	}

	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, NewHTTPError("read body failed", req, resp, err)
	}

	var b Body
	if err := json.Unmarshal(bytes, &b); err != nil {
		return nil, false, NewHTTPError("json failed", req, resp, err)
	}
	return &b, true, nil
}
