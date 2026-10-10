package auth

import (
	"github.com/AniketDubey199/models"
)

func Register(user models.User) error {
	err := Validate(user)
	if err != nil {
		return err
	}

	return nil
}
