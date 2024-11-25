package models

import (
	"time"

	"gorm.io/gorm"
)

// Patient represents a patient in the system
// @Description Patient represents a patient in the system
type Patient struct {
	gorm.Model
	FirstName          string    `json:"firstName" form:"FirstName"`
	LastName           string    `json:"lastName,omitempty" form:"LastName"`
	PhoneNumber        string    `json:"phone_number" form:"PhoneNumber"`
	Email              string    `json:"email" form:"Email"`
	BloodType          string    `json:"bloodType" form:"BloodType"`
	BirthDate          time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender             string    `json:"gender" form:"Gender"`
	Address            string    `json:"address" form:"Address"`
	CarAvailable       bool      `json:"car_available" form:"CarAvailable"`
	Urgency            int       `json:"urgency" form:"Urgency"`                         // Requests
	RedRequired        int       `json:"red-required" form:"RedRequired"`                // Requests
	RedReceived        int       `json:"red-recieved" form:"RedReceived"`                // Requests
	PlateletRequired   int       `json:"platelet-required" form:"PlateletRequired"`      // Requests
	PlateletReceived   int       `json:"platelet-recieved" form:"PlateletReceived"`      // Requests
	HospitalName       string    `json:"hospital_name" form:"HospitalName"`              // Requests
	Relationship       string    `json:"relationship" form:"Relationship"`               // Requests
	AcceptsRedCrescent bool      `json:"accepts_red_crescent" form:"AcceptsRedCrescent"` // Requests
	RedCrescentCode    string    `json:"red_crescent_code" form:"RedCrescentCode"`       // Requests
	DeathStatus        bool      `json:"death_status" form:"DeathStatus"`
	DeathDate          time.Time `json:"death_date,omitempty" form:"DeathDate" time_format:"2006-01-02"`
	SpecialPatient     bool      `json:"special_patient" form:"SpecialPatient"`
	Verify             string    `json:"verify,omitempty"`
}
