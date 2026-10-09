package auth

import (
	"errors"
	"net/mail"

	"github.com/AniketDubey199/models"
)

func Register(user models.User) error {

	if user.UserName == "" || user.Email == "" || user.Password == "" || user.Phone == "" {
		return errors.New("Fields cannot be empty")
	}
	if len(user.UserName) > 10 {
		return errors.New("Too long for username")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("Invalid email type")
	}
	return nil
}
