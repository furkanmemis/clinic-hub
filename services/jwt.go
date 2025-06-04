package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("furkanzeynep")

func GenerateJWT(email string, tenantId string) (string, error) {
	claims := jwt.MapClaims{
		"email":    email,
		"tenantId": tenantId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
