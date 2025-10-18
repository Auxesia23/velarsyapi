package dto

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}
