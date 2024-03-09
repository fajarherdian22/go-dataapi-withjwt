package service

import (
	"context"
	"database/sql"
	"gojwt/helper"
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

func (service *AuthServiceImpl) GetUserByUsername(ctx context.Context, UserName string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.IsError(err)
	defer helper.CommitOrRollback(tx)

	auths, err := service.AuthRepository.GetUserByUsername(ctx, tx, UserName)
	helper.IsError(err)

	return helper.ToUserResponse(auths)
}
