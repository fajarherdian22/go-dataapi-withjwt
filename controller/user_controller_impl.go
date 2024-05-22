package controller

import (
	"gojwt/helper"
	"gojwt/model/domain"
	"gojwt/model/web"
	"gojwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController interface {
	ValidateUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
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
		c.Abort()
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

	tokenString, err := createToken(u.Username, userResponse.Auths)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Data:   err.Error(),
			Status: false,
		})
		return

	}
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, web.WebResponse{
		Code: http.StatusOK,
		Data: gin.H{
			"Token": tokenString,
		},
		Status: true,
	})

}

func (controller *UserControllerImpl) GetAllUser(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No claims found"})
		return
	}

	jwtClaims := claims.(jwt.MapClaims)

	level := (jwtClaims)["level"]

	var WebResponse web.WebResponse
	switch level {
	case "admin":
		WebResponse = web.WebResponse{
			Code:   200,
			Data:   controller.AuthService.GetAllUser(c.Request.Context()),
			Status: true,
		}
	case "user":
		WebResponse = web.WebResponse{
			Code:   http.StatusUnauthorized,
			Data:   "don't have permission",
			Status: false,
		}
	default:
		WebResponse = web.WebResponse{
			Code:   http.StatusUnauthorized,
			Data:   "Unauthorized",
			Status: false,
		}
	}
	helper.HandleEncodeWriteJson(c, WebResponse)
}

func (controller *UserControllerImpl) CreateUser(c *gin.Context) {
	createUserRequest := web.CreateUser{}
	helper.HandleDecodeReqJson(c, &createUserRequest)

	createUserResponse := controller.AuthService.CreateUser(c.Request.Context(), createUserRequest)
	WebResponse := web.WebResponse{
		Code:   200,
		Data:   createUserResponse,
		Status: true,
	}

	helper.HandleEncodeWriteJson(c, WebResponse)
}
