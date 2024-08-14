package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

const secretKey = "secretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1800).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse the access token: " + err.Error())
	}

	if !parsedToken.Valid {
		return 0, errors.New("the access token provided is not valid")
	}

	// Extract claims
	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	userId := int64(claims["userId"].(float64))
	if !ok {
		return 0, errors.New("userId claim is missing or invalid")
	}

	return userId, nil
}
