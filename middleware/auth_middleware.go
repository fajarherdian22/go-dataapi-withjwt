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

		tokenCookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			return
		}

		// tokenString := c.Request.Header.Get("Authorization")
		tokenString := tokenCookie
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Data:   "Unauthorized to logins",
				Status: false,
			})
			c.Abort()
			return
		}
		// if !strings.HasPrefix(tokenString, "Bearer ") {
		// 	c.JSON(http.StatusUnauthorized, web.WebResponse{
		// 		Code:   http.StatusUnauthorized,
		// 		Data:   "Invalid Token format",
		// 		Status: false,
		// 	})
		// 	c.Abort()
		// 	return
		// }
		if tokenString != tokenString {
			c.JSON(http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Data:   "Invalid Token format",
				Status: false,
			})
			c.Abort()
			return
		}
		// tokenString = tokenString[len("Bearer "):]

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

		fmt.Println("this is claim : /n", token.Claims)
		c.Set("claims", token.Claims)
		c.Next()
	}
}
