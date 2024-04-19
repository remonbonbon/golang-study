package controller

import (
	"encoding/json"
	"net/http"

	"example.com/golang-study/config"
	// "github.com/go-chi/httplog/v2"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	// logger := httplog.LogEntry(r.Context())
	conf := config.Get()

	j, err := json.Marshal(conf)
	if err != nil {
		WriteError(w, r, BadRequest("エラーですよ", &err))
		return
	}

	w.Write([]byte(j))
}
