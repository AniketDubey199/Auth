package auth

import (
	"errors"

	"github.com/AniketDubey199/models"
)

func Register(user models.User) error {

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
