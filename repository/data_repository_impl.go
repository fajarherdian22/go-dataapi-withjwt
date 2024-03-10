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
)

type DataRepositoryImpl struct {
}

func NewDataRepository() DataRepository {
	return &DataRepositoryImpl{}
}

func (repository *DataRepositoryImpl) GetDataByFilter(ctx context.Context, tx *sql.Tx, level string, name string) ([]domain.Data4G, error) {
	var datas []domain.Data4G
	Query := model.Query4G_Filter
	rows, err := tx.QueryContext(ctx, fmt.Sprintf(Query, level, name))
	helper.IsError(err)
	defer rows.Close()

	if rows.Next() {
		data := domain.Data4G{}
		columns := helper.ScanDB(data)
		err := rows.Scan(columns...)
		helper.IsError(err)
		datas = append(datas, data)
		return datas, nil
	} else {
		return datas, errors.New("category is not found")
	}
}

func (repository *DataRepositoryImpl) GetAllData(ctx context.Context, tx *sql.Tx) []domain.Data4G {
	var datas []domain.Data4G
	Query := model.Query4G
	rows, err := tx.QueryContext(ctx, Query)
	helper.IsError(err)
	defer rows.Close()

	for rows.Next() {
		datamodel := domain.Data4G{}
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

func (repository *DataRepositoryImpl) GetAllFilter(ctx context.Context, tx *sql.Tx) []domain.FilterData4G {
	var datas []domain.FilterData4G
	Query := model.Query4G_Level
	rows, err := tx.QueryContext(ctx, Query)
	helper.IsError(err)
	defer rows.Close()

	for rows.Next() {
		datamodel := domain.FilterData4G{}
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
