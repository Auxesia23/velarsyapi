package utils

import (
	"errors"
	"os"
	"time"

	"github.com/Auxesia23/velarsyapi/internal/dto"
	"github.com/Auxesia23/velarsyapi/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User) string {
	var secret []byte
	secret = []byte(os.Getenv("JWT_SECRET"))

	claims := dto.UserClaims{
		User: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:   "velarsyapi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ValidateToken(tokenString string) (*dto.UserClaims, error) {
	var secret []byte
	secret = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &dto.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Failed to validate token")
}
