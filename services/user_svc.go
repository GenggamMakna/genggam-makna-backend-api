package services

import (
	"genggam-makna-api/dto"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *compServices) RegisterUserCredential(data dto.User) (*string, error) {
	uuid, err := s.repo.RegisterUserCredential(data)
	if err != nil {
		return nil, err
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = uuid
	claims["email"] = data.Email
	claims["first_name"] = data.FirstName
	claims["last_name"] = data.LastName

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
