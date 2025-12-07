package api

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "txt")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Welcome to your personal library!"))
}
