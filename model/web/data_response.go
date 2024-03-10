package web

type Data4G_Response struct {
	Short_name                         *string  `json:"short_name"`
	Hour                               *string  `json:"hour"`
	Date                               *string  `json:"date"`
	Kecamatan                          *string  `json:"kecamatan"`
	Vbt_micro_cluster                  *string  `json:"vbt_micro_cluster"`
	Vbt_sales_area                     *string  `json:"vbt_sales_area"`
	Rpt_region                         *string  `json:"rpt_region"`
	Rpt_area                           *string  `json:"rpt_area"`
	Ioh_active_user_hourly             *float64 `json:"ioh_active_user_hourly"`
	Ioh_cell_availability_nom_hourly   *float64 `json:"ioh_cell_availability_nom_hourly"`
	Ioh_cell_availability_denom_hourly *float64 `json:"ioh_cell_availability_denom_hourly"`
	Ioh_data_traffic_hourly            *float64 `json:"ioh_data_traffic_hourly"`
	Ioh_volte_traffic_hourly           *float64 `json:"ioh_volte_traffic_hourly"`
}

type FilterData4G_Response struct {
	Rpt_region        *string `json:"rpt_region"`
	Rpt_area          *string `json:"rpt_area"`
	Vbt_micro_cluster *string `json:"vbt_micro_cluster"`
	Kecamatan         *string `json:"kecamatan"`
}
