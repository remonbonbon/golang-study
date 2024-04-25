package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, r *http.Request, v any) {
	JsonWithStatus(w, r, v, http.StatusOK)
}

func JsonWithStatus(w http.ResponseWriter, r *http.Request, v any, status int) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Write(bytes)
}
