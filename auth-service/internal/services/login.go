// business logic
package services

import (
	"errors"

	usersModels "auth-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) LoginUser(ul usersModels.UserLogin) (string, string, error) {
	if ul.Email == "" || ul.Password == "" {
		return "", "", errors.New("empty email or password")
	}

	if err := s.userRepo.CheckPasswordHash(ul); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", errors.New("invalid password or email")
		}
		return "", "", errors.New("failed to check match passwords")
	}

	uid, err := s.userRepo.GetUserID(ul)
	if err != nil {
		return "", "", errors.New("failed to get UserID")
	}

	token, err := s.tokenGen.CreateJwt(uid)
	if err != nil {
		return "", "", errors.New("failed to create JWT")
	}
	return token, uid, nil
}
