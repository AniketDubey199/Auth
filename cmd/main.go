package main

import (
	"net/http"

	"github.com/AniketDubey199/internal/auth"
)

func main() {
	mux := auth.Routes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	http.ListenAndServe(server.Addr, mux)
}
