package main

import (
	"gojwt/app"
	"gojwt/controller"
	"gojwt/helper"
	"gojwt/middleware"
	"gojwt/model"
	"gojwt/repository"
	"os"

	"gojwt/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

func main() {
	var config model.Config
	configdb, err := os.ReadFile("configs.yaml")
	helper.IsError(err)

	err = yaml.Unmarshal(configdb, &config)
	helper.IsError(err)

	config_credent := config.Databases["db_credential"]
	config_data := config.Databases["db_data"]

	db_credential := app.ConDB(config_credent)
	db_data := app.ConDB(config_data)

	validate := validator.New()

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, db_credential, validate)
	authController := controller.NewUserController(authService)

	dataRepository := repository.NewDataRepository()
	dataService := service.NewDataService(dataRepository, db_data, validate)
	dataController := controller.NewDataController(dataService)

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	r := router.Group("/api/")

	{
		r.POST("/login", authController.ValidateUser)

		auth := r.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.GET("/data", dataController.GetDataAll)
			auth.GET("/data/list", dataController.GetAllFilter)
			auth.POST("/data/filter", dataController.GetDataByFilter)
			auth.GET("/alluser", authController.GetAllUser)
		}

		r.POST("/signup", authController.CreateUser)

	}

	router.Run(":8000")
}
