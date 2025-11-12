package services

import "user-service/internal/models"

type UserRepository interface {
	AddUserIntoDB(u models.User) error
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
