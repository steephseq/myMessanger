package services

import (
	"auth-service/internal/models"
	"context"
	"errors"
	"pkg/proto/userpb"
)

func (s *AuthService) RegisterUser(ur models.UserRegister) error {
	if ur.Email == "" || ur.Password == "" || ur.Name == "" || ur.Username == "" {
		return errors.New("empty email or password")
	}

	exists, err := s.userRepo.ExistsUser(ur)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}

	if err := s.userRepo.CreateUser(ur); err != nil {
		return err
	}

	_, err = s.userClient.CreateUser(context.Background(), &userpb.CreateUserRequest{
		Email:    ur.Email,
		Password: ur.Password,
		Name:     ur.Name,
		Username: ur.Username,
	})
	if err != nil {
		return err
	}
	return nil
}
