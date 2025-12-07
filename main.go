package main

import (
	"log"
	"net/http"

	"github.com/gopher-95/library_porject/pkg/api"
)

func main() {
	api.Init()

	log.Println("Server is running on port 8080!")
	err := http.ListenAndServe(":8080", api.Mux)
	if err != nil {
		panic(err)
	}
}
