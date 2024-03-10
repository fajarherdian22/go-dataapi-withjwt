package repository

import (
	"context"
	"database/sql"
	"gojwt/model/domain"
)

type DataRepository interface {
	GetDataByFilter(ctx context.Context, tx *sql.Tx, level, name string) ([]domain.Data4G, error)
	GetAllData(ctx context.Context, tx *sql.Tx) []domain.Data4G
	GetAllFilter(ctx context.Context, tx *sql.Tx) []domain.FilterData4G
}
