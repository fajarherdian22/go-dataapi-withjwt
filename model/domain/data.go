package domain

type Data4G struct {
	Short_name                         *string
	Hour                               *string
	Date                               *string
	Kecamatan                          *string
	Vbt_micro_cluster                  *string
	Vbt_sales_area                     *string
	Rpt_region                         *string
	Rpt_area                           *string
	Ioh_active_user_hourly             *float64
	Ioh_cell_availability_nom_hourly   *float64
	Ioh_cell_availability_denom_hourly *float64
	Ioh_data_traffic_hourly            *float64
	Ioh_volte_traffic_hourly           *float64
}
