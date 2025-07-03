package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Dummy secretKey in production it should be created in a more secure manner
const secretKey = "supersecret"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("could not pass token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}
	/*
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("invalid token claims")
		}

		email := claims["email"].(string)
		userID := claims["userID"].(int64) */
	return nil
}
