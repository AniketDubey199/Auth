package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokens(username string) (string, error) {
	key := os.Getenv("SECRET_KEY")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"expires_at": time.Now().Add(24 * time.Hour),
	})
	tokens, err := t.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("Error in generating singed integer")
	}
	return tokens, nil
}
