package main

import (
	"log"
	"net/http"

	"github.com/AniketDubey199/database"
	"github.com/AniketDubey199/internal/auth"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	service := auth.NewService(db)
	mux := auth.Routes(service)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	http.ListenAndServe(server.Addr, mux)
}
