package middleware

import (
	"fmt"
	"gojwt/controller"
	"gojwt/model/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return controller.SECRET_KEY, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Invalids token")
	}

	return nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Data:   "Unauthorized to login",
				Status: false,
			})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Data:   "Invalid Token",
				Status: false,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
