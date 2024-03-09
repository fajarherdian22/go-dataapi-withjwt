package service

import (
	"context"
	"database/sql"
	"gojwt/helper"
	"gojwt/model/web"
	"gojwt/repository"

	"github.com/go-playground/validator/v10"
)

type DataServiceImpl struct {
	DataRepository repository.DataRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewDataService(DataRepository repository.DataRepository, DB *sql.DB, Validate *validator.Validate) DataService {
	return &DataServiceImpl{
		DataRepository: DataRepository,
		DB:             DB,
		Validate:       Validate,
	}
}

func (service *DataServiceImpl) GetDataByFilter(ctx context.Context, level string, name string) []web.Data4G_Response {
	tx, err := service.DB.Begin()
	helper.IsError(err)
	defer helper.CommitOrRollback(tx)

	data, err := service.DataRepository.GetDataByFilter(ctx, tx, level, name)
	helper.IsError(err)

	return helper.ToDataResponses(data)
}

func (service *DataServiceImpl) GetAllData(ctx context.Context) []web.Data4G_Response {
	tx, err := service.DB.Begin()
	helper.IsError(err)
	defer helper.CommitOrRollback(tx)

	data := service.DataRepository.GetAllData(ctx, tx)

	return helper.ToDataResponses(data)
}
