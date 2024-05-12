package helper

import jwt "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Level string `json:"level"`
	jwt.StandardClaims
}

var SECRET_KEY = []byte("secretkeys")

func ParseJwt(cookie string) (string, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	return claims.Level, nil
}
