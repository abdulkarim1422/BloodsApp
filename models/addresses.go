package models

import "gorm.io/gorm"

// Countries represents a country in the system
// @Description Countries represents a country in the system
type Countries struct {
	gorm.Model
	CodeNumber  string `json:"code_number" form:"CodeNumber"`
	CountryName string `json:"country_name" form:"CountryName"`
}

// Cities represents a city in the system
// @Description Cities represents a city in the system
type Cities struct {
	gorm.Model
	CodeNumber string `json:"code_number" form:"CodeNumber"`
	CityName   string `json:"city_name" form:"CityName"`
	CountryID  int    `json:"country_id" form:"CountryID"`
}

// District represents a district in the system
// @Description District represents a district in the system
type District struct {
	gorm.Model
	DistrictName string `json:"district_name" form:"DistrictName"`
	CityID       int    `json:"city_id" form:"CityID"`
}

// Hospitals represents a hospital in the system
// @Description Hospitals represents a hospital in the system
type Hospitals struct {
	gorm.Model
	HospitalName string `json:"hospital_name" form:"HospitalName"`
	DistrictID   int    `json:"district_id" form:"DistrictID"`
	PostalCode   int    `json:"postal_code" form:"PostalCode"`
}
