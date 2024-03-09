package repository

import (
	"context"
	"database/sql"
	"gojwt/model/domain"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, tx *sql.Tx, UserName string) (domain.User, error)
}
