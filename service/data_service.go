package service

import (
	"context"
	"gojwt/model/web"
)

type DataService interface {
	GetDataByFilter(ctx context.Context, level, name string) []web.Data4G_Response
	GetAllData(ctx context.Context) []web.Data4G_Response
	GetAllFilter(ctx context.Context) []web.FilterData4G_Response
}
