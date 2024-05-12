package service

import (
	"context"
	"database/sql"
	"fmt"
	"gojwt/helper"
	"gojwt/model/domain"
	"gojwt/model/web"
	"gojwt/repository"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(AuthRepository repository.AuthRepository, DB *sql.DB, Validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: AuthRepository,
		DB:             DB,
		Validate:       Validate,
	}
}

func (service *AuthServiceImpl) GetAllUser(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.IsError(err)
	defer helper.CommitOrRollback(tx)

	data := service.AuthRepository.GetAllUser(ctx, tx)

	return helper.ToUserResponses(data)
}

func (service *AuthServiceImpl) CreateUser(ctx context.Context, payload web.CreateUser) web.UserResponse {

	fmt.Println("ini service")
	err := service.Validate.Struct(payload)
	helper.IsError(err)

	tx, err := service.DB.Begin()
	helper.IsError(err)
	defer helper.CommitOrRollback(tx)

	PayloadCreate := domain.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}

	fmt.Println("ini payload \n", payload)

	PayloadCreate = service.AuthRepository.CreateUser(ctx, tx, PayloadCreate)

	return helper.ToUserResponse(PayloadCreate)

}

func (service *AuthServiceImpl) GetUserByUsername(ctx context.Context, UserName string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.IsError(err)
	defer helper.CommitOrRollback(tx)

	auths, err := service.AuthRepository.GetUserByUsername(ctx, tx, UserName)
	helper.IsError(err)

	return helper.ToUserResponse(auths)
}
