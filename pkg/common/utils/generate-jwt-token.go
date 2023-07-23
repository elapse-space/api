package utils

import (
	"api/pkg/common/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWTToken(id int32) (string, error) {
	c, err := config.LoadConfig()
	var jwtSecret = []byte(c.JWTSecret)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
