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
