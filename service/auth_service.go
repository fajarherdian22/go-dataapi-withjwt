package service

import (
	"context"
	"gojwt/model/web"
)

type AuthService interface {
	GetUserByUsername(ctx context.Context, UserName string) web.UserResponse
}
