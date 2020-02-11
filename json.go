package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func JSON(fn handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		i, err := fn(w, r)

		if err != nil {
			if IsClientError(err) {
				Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if i == nil {
			return
		}
		_ = json.NewEncoder(w).Encode(i)
	}
}

func Error(w http.ResponseWriter, err string, code int) {
	w.Header().Add("Content-Type", "application/json")
	log.Println(code, ":", err)

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	}{
		Error:      err,
		StatusCode: code,
	})
}
