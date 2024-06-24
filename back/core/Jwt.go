package core

import (
	"github.com/dgrijalva/jwt-go"
)

func JwtGenerator(u User, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.Id,
		"name":  u.Name,
		"email": u.Email,
	})

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return err.Error()
	}

	return tokenString
}
