package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gojwt/helper"
	"gojwt/model"
	"gojwt/model/domain"
	"reflect"

	"github.com/google/uuid"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetAllUser(ctx context.Context, tx *sql.Tx) []domain.User {
	var datas []domain.User
	Query := model.QueryGetAllUser
	rows, err := tx.QueryContext(ctx, Query)
	helper.IsError(err)
	defer rows.Close()

	for rows.Next() {
		datamodel := domain.User{}
		s := reflect.ValueOf(&datamodel).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := rows.Scan(columns...)
		helper.IsError(err)
		datas = append(datas, datamodel)
	}
	return datas
}

func (repository *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sql.Tx, UserName string) (domain.User, error) {
	Query := model.QueryGetUser
	rows, err := tx.QueryContext(ctx, fmt.Sprintf(Query, UserName))
	helper.IsError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Auths)
		helper.IsError(err)
		return user, nil
	} else {
		return user, errors.New("username is not found")
	}

}

func (repository *AuthRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, payload domain.User) domain.User {
	Query := model.QueryCreateUser
	unique_id := uuid.New()
	_, err := tx.ExecContext(ctx, fmt.Sprintf(Query, unique_id, payload.Username, payload.Password, payload.Email))
	helper.IsError(err)

	return payload

}
