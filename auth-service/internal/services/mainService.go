package services

import (
	"auth-service/internal/clients"
	"auth-service/internal/models"
)

type UserRepository interface {
	CheckPasswordHash(ul models.UserLogin) error
	GetUserID(ul models.UserLogin) (string, error)
	CreateUser(ul models.UserRegister) error
	ExistsUser(ul models.UserRegister) (bool, error)
}

type TokenGenerator interface {
	CreateJwt(uid string) (string, error)
}

type AuthService struct {
	userRepo   UserRepository
	tokenGen   TokenGenerator
	userClient *clients.UserClient
}

func NewAuthService(userRepo UserRepository, tokenGen TokenGenerator, userClient *clients.UserClient) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		tokenGen:   tokenGen,
		userClient: userClient,
	}
}
