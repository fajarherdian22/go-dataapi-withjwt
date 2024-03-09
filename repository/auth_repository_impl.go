package repository

import (
	"context"
	"database/sql"
	"errors"
	"gojwt/helper"
	"gojwt/model/domain"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sql.Tx, UserName string) (domain.User, error) {
	SQL := "select id, username, password, email from credential where username = ?"
	rows, err := tx.QueryContext(ctx, SQL, UserName)
	helper.IsError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		helper.IsError(err)
		return user, nil
	} else {
		return user, errors.New("usersname is not found")
	}

}

// func (repository *AuthRepositoryImpl) GetUserAll(ctx context.Context, tx *sql.Tx) []domain.User {
// 	SQL := "select id, username, password, email from credential"
// 	rows, err := tx.QueryContext(ctx, SQL)
// 	helper.IsError(err)
// 	defer rows.Close()

// 	var users []domain.User
// 	for rows.Next() {
// 		user := domain.User{}
// 		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
// 		helper.IsError(err)
// 		users = append(users, user)
// 	}
// 	return users
// }
