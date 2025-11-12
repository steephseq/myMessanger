package utils

import (
	"auth-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordFunc(u models.UserRegister) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}
