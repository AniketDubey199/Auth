package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connection() (*sql.DB, error) {
	pass := os.Getenv("POSTGRES_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", pass, "Aniket")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	if err != nil {
		defer db.Close()
		log.Panic(err)
	}
	createtable := `create table if not exists Users (
		id serial primary key,
		name text not null ,
		email text not null unique, 
		password text not null,
		created_at timestamptz not null default now()
	)`
	_, err = db.Exec(createtable)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
