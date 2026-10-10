package main

import (
	"log"
	"net/http"

	"github.com/AniketDubey199/database"
	"github.com/AniketDubey199/internal/auth"
	"github.com/joho/godotenv"
)

func main() {
	mux := auth.Routes()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	database.Connection()

	http.ListenAndServe(server.Addr, mux)
}
