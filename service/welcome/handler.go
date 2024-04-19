package welcome

import (
	"encoding/json"
	"net/http"

	"example.com/golang-study/config"
	srv "example.com/golang-study/service"
	// "github.com/go-chi/httplog/v2"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// logger := httplog.LogEntry(r.Context())
	conf := config.Get()

	j, err := json.Marshal(conf)
	if err != nil {
		srv.WriteError(w, r, srv.BadRequest("エラーですよ", &err))
		return
	}

	w.Write([]byte(j))
}
