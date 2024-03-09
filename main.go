package main

import (
	"fmt"
	"gojwt/app"
	"gojwt/controller"
	"gojwt/middleware"
	"gojwt/repository"

	"gojwt/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	db_credential := app.CredentialDB()
	db_data := app.DataDB()
	validate := validator.New()
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db_credential, validate)
	authController := controller.NewUserController(authService)

	dataRepository := repository.NewDataRepository()
	dataService := service.NewDataService(dataRepository, db_data, validate)
	dataController := controller.NewDataController(dataService)

	router := gin.New()
	router.Use(gin.Recovery())
	// router.Use(middleware.AuthMiddleware())

	router.POST("/login", authController.ValidateUser)
	router.GET("/authenticate", middleware.AuthMiddleware())

	router.GET("/data", dataController.GetDataAll)
	router.POST("/datafilter", dataController.GetDataByFilter)

	router.Run(":8000")

	fmt.Println("Server is running")

}
