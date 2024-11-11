package services

import (
	"time"
)

type Donor struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name" form:"FirstName"`
	LastName      string    `json:"last_name,omitempty" form:"LastName"`
	PhoneNumber   string    `json:"phone_number" form:"PhoneNumber"`
	BloodType     string    `json:"bloodType" form:"BloodType"`
	BirthDate     time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender        string    `json:"gender" form:"Gender"`
	Address       string    `json:"address" form:"Address"`
	City          string    `json:"city" form:"City"`
	CarAvailable  bool      `json:"car_available,omitempty" form:"CarAvailable"`
	Paid          bool      `json:"paid,omitempty"`
	RedTimer      time.Time `json:"red_timer"`
	PlateletTimer time.Time `json:"platelet_timer"`
	Score         int       `json:"score"`
	Verify        string    `json:"verify,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Patient struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName" form:"FirstName"`
	LastName         string    `json:"lastName,omitempty" form:"LastName"`
	PhoneNumber      string    `json:"phone_number" form:"PhoneNumber"`
	BloodType        string    `json:"bloodType" form:"BloodType"`
	BirthDate        time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender           string    `json:"gender" form:"Gender"`
	Address          string    `json:"address" form:"Address"`
	CarAvailable     bool      `json:"car_available" form:"CarAvailable"`
	Urgency          int       `json:"urgency" form:"Urgency"`                    // Requests
	RedRequired      int       `json:"red-required" form:"RedRequired"`           // Requests
	RedReceived      int       `json:"red-recieved" form:"RedReceived"`           // Requests
	PlateletRequired int       `json:"platelet-required" form:"PlateletRequired"` // Requests
	PlateletReceived int       `json:"platelet-recieved" form:"PlateletReceived"` // Requests
	HospitalName     string    `json:"hospital_name" form:"HospitalName"`         // Requests
	Verify           string    `json:"verify,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Requests struct {
	ID               int    `json:"id"`
	PatientID        int    `json:"patient_id"`
	DonorID          int    `json:"donor_id,omitempty"`
	Red              int    `json:"red" form:"Red"`
	Platelet         int    `json:"platelet" form:"Platelet"`
	CarAvailable     bool   `json:"car_available" form:"CarAvailable"`
	RedRequired      int    `json:"red-required" form:"RedRequired"`           // Requests
	RedReceived      int    `json:"red-recieved" form:"RedReceived"`           // Requests
	PlateletRequired int    `json:"platelet-required" form:"PlateletRequired"` // Requests
	PlateletReceived int    `json:"platelet-recieved" form:"PlateletReceived"` // Requests
	HospitalName     string `json:"hospital_name" form:"HospitalName"`         // Requests
	RequestStatus    int    `json:"request_status" form:"RequestStatus"`
	Urgency          int    `json:"urgency" form:"Urgency"` // Requests
	Relationship     int    `json:"relationship" form:"Relationship"`
	//
}

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

var donors = []Donor{
	{ID: 1, FirstName: "Donor One", LastName: "DDD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 2, FirstName: "Donor Two", LastName: "DDD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 3, FirstName: "Donor Three", LastName: "DDD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 4, FirstName: "Donor Four", LastName: "DDD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 5, FirstName: "Donor Five", LastName: "DDD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 6, FirstName: "Donor Six", LastName: "DDD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 7, FirstName: "Donor Seven", LastName: "DDD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 8, FirstName: "Donor Eight", LastName: "DDD", BloodType: "B-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 9, FirstName: "Donor Nine", LastName: "DDD", BloodType: "A-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 10, FirstName: "Donor Ten", LastName: "DDD", BloodType: "AB-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 11, FirstName: "Donor Eleven", LastName: "DDD", BloodType: "AB-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 12, FirstName: "Donor Twelve", LastName: "DDD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "+905511678290", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
}

var patients = []Patient{
	{ID: 1, FirstName: "Patient One", LastName: "PPP", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 1, RedReceived: 0, PlateletRequired: 1, PlateletReceived: 0, CarAvailable: true, HospitalName: "Hospital One"},
	{ID: 2, FirstName: "Patient Two", LastName: "PPP", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 2, RedReceived: 0, PlateletRequired: 0, PlateletReceived: 0, CarAvailable: false, HospitalName: "Hospital Two"},
	{ID: 3, FirstName: "Patient Three", LastName: "PPP", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 1, RedReceived: 0, PlateletRequired: 1, PlateletReceived: 0, CarAvailable: true, HospitalName: "Hospital Three"},
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
