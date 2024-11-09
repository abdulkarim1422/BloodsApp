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
	CarAvailable  bool      `json:"car_available,omitempty" form:"CarAvailable"`
	Paid          bool      `json:"paid,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Address       string    `json:"address" form:"Address"`
	RedTimer      time.Time `json:"red_timer"`
	PlateletTimer time.Time `json:"platelet_timer"`
	Score         int       `json:"score"`
	Verify        string    `json:"verify,omitempty"`
}

type Patient struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName" form:"FirstName"`
	LastName         string    `json:"lastName,omitempty" form:"LastName"`
	BloodType        string    `json:"bloodType" form:"BloodType"`
	BirthDate        time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender           string    `json:"gender" form:"Gender"`
	Address          string    `json:"address" form:"Address"`
	PhoneNumber      string    `json:"phone" form:"Phone"`
	Urgency          int       `json:"urgency" form:"Urgency"`
	RedQuantity      int       `json:"red-quantity" form:"RedQuantity"`
	PlateletQuantity int       `json:"platelet-quantity" form:"PlateletQuantity"`
	CarAvailable     bool      `json:"car_available" form:"CarAvailable"`
	HospitalName     string    `json:"hospital_name" form:"HospitalName"`
	Verify           string    `json:"verify,omitempty"`
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
	{ID: 1, FirstName: "Patient One", LastName: "PPP", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, CarAvailable: true, HospitalName: "Hospital One"},
	{ID: 2, FirstName: "Patient Two", LastName: "PPP", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedQuantity: 2, PlateletQuantity: 0, CarAvailable: false, HospitalName: "Hospital Two"},
	{ID: 3, FirstName: "Patient Three", LastName: "PPP", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, CarAvailable: true, HospitalName: "Hospital Three"},
}