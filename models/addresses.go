package models

import "gorm.io/gorm"

type Countries struct {
	gorm.Model
	CodeNumber  string `json:"code_number" form:"CodeNumber"`
	CountryName string `json:"country_name" form:"CountryName"`
}

type Cities struct {
	gorm.Model
	CodeNumber string `json:"code_number" form:"CodeNumber"`
	CityName   string `json:"city_name" form:"CityName"`
	CountryID  int    `json:"country_id" form:"CountryID"`
}

type District struct {
	gorm.Model
	DistrictName string `json:"district_name" form:"DistrictName"`
	CityID       int    `json:"city_id" form:"CityID"`
}

type Hospitals struct {
	gorm.Model
	HospitalName string `json:"hospital_name" form:"HospitalName"`
	DistrictID   int    `json:"district_id" form:"DistrictID"`
	PostalCode   int    `json:"postal_code" form:"PostalCode"`
}
