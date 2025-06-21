package utils

import (
	"backend/config"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	secret := config.Env.JWT_SECRET
	// IMPORTANT: Add a check to ensure the secret is not empty
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable not set or empty")
	}

	// --- THIS IS THE CRUCIAL CHANGE ---
	t, err := token.SignedString([]byte(secret)) // Convert the string secret to []byte
	// ----------------------------------

	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
