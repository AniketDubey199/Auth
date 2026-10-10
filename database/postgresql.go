package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connection() {
	pass := os.Getenv("POSTGRES_PASSWORD")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", pass, "Aniket")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Successfully connected!")
}
