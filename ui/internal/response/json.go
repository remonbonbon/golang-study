package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, r *http.Request, v any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
