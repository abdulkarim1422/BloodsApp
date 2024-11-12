package models

type Countries struct {
	ID          int    `json:"id"`
	CodeNumber  string `json:"code_number" form:"CodeNumber"`
	CountryName string `json:"country_name" form:"CountryName"`
}

type Cities struct {
	ID         int    `json:"id"`
	CodeNumber string `json:"code_number" form:"CodeNumber"`
	CityName   string `json:"city_name" form:"CityName"`
	CountryID  int    `json:"country_id" form:"CountryID"`
}

type District struct {
	ID           int    `json:"id"`
	DistrictName string `json:"district_name" form:"DistrictName"`
	CityID       int    `json:"city_id" form:"CityID"`
}

type Hospitals struct {
	ID           int    `json:"id"`
	HospitalName string `json:"hospital_name" form:"HospitalName"`
	DistrictID   int    `json:"district_id" form:"DistrictID"`
	PostalCode   int    `json:"postal_code" form:"PostalCode"`
}

var CountriesList = []Countries{
	{ID: 1, CountryName: "Turkey"},
}

var CitiesList = []Cities{
	{ID: 1, CodeNumber: "TR34", CityName: "Istanbul", CountryID: 1},
	{ID: 2, CodeNumber: "TR50", CityName: "Ankara", CountryID: 1},
	{ID: 3, CodeNumber: "TR51", CityName: "Izmir", CountryID: 1},
}

var DistrictList = []District{
	{ID: 1, DistrictName: "Kadikoy", CityID: 1},
	{ID: 2, DistrictName: "Besiktas", CityID: 1},
	{ID: 3, DistrictName: "Cankaya", CityID: 2},
	{ID: 4, DistrictName: "Bornova", CityID: 3},
}

var HospitalsList = []Hospitals{
	{ID: 1, HospitalName: "Acibadem", DistrictID: 1, PostalCode: 34710},
	{ID: 2, HospitalName: "Memorial", DistrictID: 2, PostalCode: 34330},
	{ID: 3, HospitalName: "Ankara Hospital", DistrictID: 3, PostalCode: 06510},
	{ID: 4, HospitalName: "Izmir Hospital", DistrictID: 4, PostalCode: 35030},
}
