package middleware

import (
	"fmt"
	"gojwt/controller"
	"gojwt/model/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return controller.SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			return
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Data:   "Unauthorized to logins",
				Status: false,
			})
			c.Abort()
			return
		}

		token, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Data:   "Invalid Token",
				Status: false,
			})
			c.Abort()
			return
		}
		c.Set("claims", token.Claims)
		c.Next()
	}
}
