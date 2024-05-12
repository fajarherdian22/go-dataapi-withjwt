package service

import (
	"context"
	"gojwt/model/web"
)

type AuthService interface {
	CreateUser(ctx context.Context, payload web.CreateUser) web.UserResponse
	GetAllUser(ctx context.Context) []web.UserResponse
	GetUserByUsername(ctx context.Context, UserName string) web.UserResponse
}
