package models

import "time"

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

var patients = []Patient{
	{ID: 1, FirstName: "Patient One", LastName: "PPP", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 1, RedReceived: 0, PlateletRequired: 1, PlateletReceived: 0, CarAvailable: true, HospitalName: "Hospital One"},
	{ID: 2, FirstName: "Patient Two", LastName: "PPP", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 2, RedReceived: 0, PlateletRequired: 0, PlateletReceived: 0, CarAvailable: false, HospitalName: "Hospital Two"},
	{ID: 3, FirstName: "Patient Three", LastName: "PPP", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 1, RedReceived: 0, PlateletRequired: 1, PlateletReceived: 0, CarAvailable: true, HospitalName: "Hospital Three"},
}
