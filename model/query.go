package model

var QueryGetAllUser = `SELECT * FROM credential`

var QueryGetUser = `SELECT * from credential where username = "%s"`

var QueryCreateUser = `
INSERT INTO credential (id, username, password, email, auths) 
VALUES("%s", "%s", "%s", "%s", "admin")`

var Query4G_Level = `
SELECT DISTINCT 
a.rpt_region, 
a.rpt_area, 
a.vbt_micro_cluster, 
a.kecamatan
FROM mycom_report.site_4g_kpi_hourly_mycom a;
`

var Query4G = `
SELECT 
Short_name,
Hour,
Date,
Kecamatan,
Vbt_micro_cluster,
Vbt_sales_area,
Rpt_region,
Rpt_area,
Ioh_active_user_hourly,
Ioh_cell_availability_nom_hourly,
Ioh_cell_availability_denom_hourly,
Ioh_data_traffic_hourly,
Ioh_volte_traffic_hourly
FROM site_4g_kpi_hourly_mycom a 
WHERE date = "2024-02-20" and Kecamatan = "Ciamis"
ORDER BY Hour`

var Query4G_Filter = `
SELECT 
Short_name,
Hour,
Date,
Kecamatan,
Vbt_micro_cluster,
Vbt_sales_area,
Rpt_region,
Rpt_area,
Ioh_active_user_hourly,
Ioh_cell_availability_nom_hourly,
Ioh_cell_availability_denom_hourly,
Ioh_data_traffic_hourly,
Ioh_volte_traffic_hourly
FROM site_4g_kpi_hourly_mycom a 
WHERE hour >= DATE_SUB(NOW(), INTERVAL 1 DAY)
AND %s = "%s"
ORDER BY Hour`
