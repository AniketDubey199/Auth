package auth

import (
	"errors"
	"log"
	"net/mail"
	"strings"
	"unicode"

	"github.com/AniketDubey199/models"
	"golang.org/x/crypto/bcrypt"
)

func Validate(user models.User) error {
	if user.UserName == "" || user.Email == "" || user.Password == "" {
		return errors.New("Fields cannot be empty")
	}
	if len(user.UserName) > 10 {
		return errors.New("Too long for username")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("Invalid email type")
	}
	if user.Password != "" {
		if len(user.Password) < 8 {
			return errors.New("Password too small")
		}
		if !strings.ContainsFunc(user.Password, unicode.IsUpper) {
			return errors.New("Invalid format")
		}
		if !strings.ContainsFunc(user.Password, unicode.IsDigit) {
			return errors.New("Invalid format")
		}
		if !strings.ContainsFunc(user.Password, unicode.IsSymbol) && !strings.ContainsFunc(user.Password, unicode.IsPunct) {
			return errors.New("Invalid format")
		}
	}
	return nil
}

func EncryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func DecryptPassword(userpassword string, actualpassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(actualpassword))
	if err != nil {
		return err
	}
	return nil
}
