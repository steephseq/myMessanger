package token

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret string
	expire time.Duration
}

func NewJWTManager() (*JWTManager, error) {
	expireHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		log.Println("failed to convert JWT_EXPIRE to int")
		return nil, err
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("JWT_SECRET is not set")
		return nil, errors.New("JWT_SECRET is not set")
	}
	return &JWTManager{
		secret: secret,
		expire: time.Duration(expireHours) * time.Hour,
	}, nil
}

func (j *JWTManager) CreateJwt(uid string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": uid,
		"exp":     time.Now().Add(j.expire).Unix(),
		"iss":     "auth-service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}
