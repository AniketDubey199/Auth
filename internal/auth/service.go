package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/mail"
	"time"

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

func (s *Service) Register(ctx context.Context, user models.User, response *models.UserReq) error {
	err := Validate(user)
	if err != nil {
		return err
	}
	createtable := `create table if not exists Users (
		id serial primary key,
		name text not null ,
		email text not null unique, 
		password text not null,
		created_at timestamptz not null default now()
	)`
	_, err = s.db.Exec(createtable)
	if err != nil {
		return err
	}
	user.Password = EncryptPassword(user.Password)
	register := `insert into users (name , email , password) values ($1 , $2 , $3 ) returning id`
	id := 0
	err = s.db.QueryRow(register, user.UserName, user.Email, user.Password).Scan(&id)
	if err != nil {
		return err
	}
	response.ID = id
	response.Email = user.UserName
	response.Created_at = time.Now().Format(time.UnixDate)
	fmt.Println("The New id inserted is: ", id)
	return nil
}

func (s *Service) Login(ctx context.Context, user *models.UserLogin, userEmail string, password string) error {
	if userEmail == "" || password == "" {
		return errors.New("email or password is empty")
	}
	if _, err := mail.ParseAddress(userEmail); err != nil {
		return errors.New("Invalid email type")
	}
	query := `select email , password from users where email = $1;`
	row := s.db.QueryRow(query, userEmail)
	userpass := ""
	err := row.Scan(&user.Email, &userpass) // change this password should not be visible
	if err != nil {
		return errors.New("Db scan not valid")
	}
	err = DecryptPassword(userpass, password)
	if err != nil {
		return errors.New("Password or username is not valid")
	}
	user.Password = password
	return nil
}
