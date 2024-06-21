package core

import (
	"github.com/dgrijalva/jwt-go"
)

func JwtGenerator(id int, name, username, email, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"name":     name,
		"username": username,
		"email":    email,
	})

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return err.Error()
	}

	return tokenString
}
