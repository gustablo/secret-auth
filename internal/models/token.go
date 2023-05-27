package models

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t Token) Generate(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signed, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", err
	}

	return signed, nil
}
