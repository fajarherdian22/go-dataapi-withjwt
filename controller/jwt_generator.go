package controller

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte("secretkeys")
var LOGIN_DURATION = time.Now().Add(time.Minute * 60).Unix()

func createToken(username, level string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"level":    level,
			"exp":      LOGIN_DURATION,
		})

	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// func RefreshToken(c *gin.Context){
// 	authHeader := c.GetHeader("Security")

// }
