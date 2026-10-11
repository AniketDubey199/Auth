package auth

import (
	"context"
	"database/sql"

	"github.com/AniketDubey199/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service { // think of it as a object of service and it helps to push db instance in service db
	return &Service{
		db: db,
	}
}

func (s *Service) Register(ctx context.Context, user models.User) error {
	err := Validate(user)
	if err != nil {
		return err
	}
	createtable := `create table if not exists Users (
		id serial primary key,
		name text not null,
		email text not null , 
		phone text not null,
		password text not null
	)`
	_, err = s.db.Exec(createtable)
	if err != nil {
		return err
	}

	return nil
}
