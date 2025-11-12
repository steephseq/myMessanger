package services

import (
	"errors"
	"user-service/internal/models"

	"github.com/badoux/checkmail"
)

func (s *UserService) CreateUser(u models.User) error {
	if u.Email == "" || u.Password == "" || u.Name == "" || u.Username == "" {
		return errors.New("empty fields")
	}

	if u.AvatarURL == "" {
		u.AvatarURL = "https://ru.pinterest.com/pin/612278511877184766/"
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return err
	}

	if err := s.userRepo.AddUserIntoDB(u); err != nil {
		return err
	}
	return nil
}
