package helper

import (
	"gojwt/model/domain"
	"gojwt/model/web"
)

func ToUserResponse(User domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       User.Id,
		Username: User.Username,
		Password: User.Password,
		Email:    User.Email,
		Auths:    User.Auths,
	}
}

func ToUserResponses(Users []domain.User) []web.UserResponse {
	var UserResponses []web.UserResponse

	for _, user := range Users {
		UserResponses = append(UserResponses, ToUserResponse(user))
	}
	return UserResponses
}

func ToDataResponse(Data domain.Data4G) web.Data4G_Response {
	return web.Data4G_Response{
		Short_name:                         Data.Short_name,
		Hour:                               Data.Hour,
		Date:                               Data.Date,
		Kecamatan:                          Data.Kecamatan,
		Vbt_micro_cluster:                  Data.Vbt_micro_cluster,
		Vbt_sales_area:                     Data.Vbt_sales_area,
		Rpt_region:                         Data.Rpt_region,
		Rpt_area:                           Data.Rpt_area,
		Ioh_active_user_hourly:             Data.Ioh_active_user_hourly,
		Ioh_cell_availability_nom_hourly:   Data.Ioh_cell_availability_nom_hourly,
		Ioh_cell_availability_denom_hourly: Data.Ioh_cell_availability_denom_hourly,
		Ioh_data_traffic_hourly:            Data.Ioh_data_traffic_hourly,
		Ioh_volte_traffic_hourly:           Data.Ioh_volte_traffic_hourly,
	}
}

func ToDataResponses(Datas []domain.Data4G) []web.Data4G_Response {
	var Data4GResponses []web.Data4G_Response

	for _, data := range Datas {
		Data4GResponses = append(Data4GResponses, ToDataResponse(data))
	}
	return Data4GResponses
}

func ToDataFilterResponse(Data domain.FilterData4G) web.FilterData4G_Response {
	return web.FilterData4G_Response{
		Rpt_region:        Data.Rpt_region,
		Rpt_area:          Data.Rpt_area,
		Vbt_micro_cluster: Data.Vbt_micro_cluster,
		Kecamatan:         Data.Kecamatan,
	}
}

func ToDataFilterResponses(Datas []domain.FilterData4G) []web.FilterData4G_Response {
	var FilterData4G []web.FilterData4G_Response

	for _, data := range Datas {
		FilterData4G = append(FilterData4G, ToDataFilterResponse(data))
	}
	return FilterData4G
}
