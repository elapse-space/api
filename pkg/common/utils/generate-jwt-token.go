package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWTToken(username string) (string, error) {
	var jwtSecret = []byte("g78asdg78ags87g89hhasdj")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
