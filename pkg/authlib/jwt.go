package authlib

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var Secret []byte

func InitSecret() error {
	Secret = []byte(os.Getenv("JWT_SECRET"))
	if string(Secret) == "" {
		return errors.New("JWT_SECRET is not set")
	}
	return nil
}

func GetUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, getSecret)
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to get claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("failed to get user_id")
	}
	return userID, nil
}

func getSecret(token *jwt.Token) (interface{}, error) {
	return Secret, nil
}
