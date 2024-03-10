package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gojwt/helper"
	"gojwt/model"
	"gojwt/model/domain"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sql.Tx, UserName string) (domain.User, error) {
	Query := model.QueryGetUser
	rows, err := tx.QueryContext(ctx, fmt.Sprintf(Query, UserName))
	helper.IsError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		helper.IsError(err)
		return user, nil
	} else {
		return user, errors.New("username is not found")
	}

}
