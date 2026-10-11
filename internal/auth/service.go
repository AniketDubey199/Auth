package auth

import (
	"database/sql"

	"github.com/AniketDubey199/models"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}
func (s *Service) Register(user models.User) error {
	err := Validate(user)
	if err != nil {
		return err
	}

	return nil
}
