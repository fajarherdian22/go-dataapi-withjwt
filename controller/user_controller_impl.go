package controller

import (
	"gojwt/model/domain"
	"gojwt/model/web"
	"gojwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	ValidateUser(c *gin.Context)
}

type UserControllerImpl struct {
	AuthService service.AuthService
}

func NewUserController(authService service.AuthService) UserController {
	return &UserControllerImpl{
		AuthService: authService,
	}
}

func (controller *UserControllerImpl) ValidateUser(c *gin.Context) {
	var u domain.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := controller.AuthService.GetUserByUsername(c.Request.Context(), u.Username)

	if userResponse.Password == "" || userResponse.Password != u.Password {
		c.JSON(http.StatusUnauthorized, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Data:   "Unauthorized to login",
			Status: false,
		})
		return
	}

	tokenString, err := createToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Data:   err.Error(),
			Status: false,
		})
		return
	}
	c.JSON(http.StatusOK, web.WebResponse{
		Code: http.StatusOK,
		Data: gin.H{
			"Token":    tokenString,
			"Username": u.Username,
		},
		Status: true,
	})
}
