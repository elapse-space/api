package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	Id int32 `json:"id"`
}
