package models

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserId int64 `json:"id"`
	jwt.StandardClaims
}