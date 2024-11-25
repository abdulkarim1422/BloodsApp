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
	LatinName          string    `json:"latinName,omitempty" form:"LatinName"` //form
	PhoneNumber        string    `json:"phone_number" form:"PhoneNumber"`      //form
	Email              string    `json:"email" form:"Email"`
	BloodType          string    `json:"bloodType" form:"BloodType"` //form
	BirthDate          time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender             string    `json:"gender" form:"Gender"`
	Address            Address   `json:"addresses,omitempty" gorm:"many2many:donor_addresses;"` //form*
	CarAvailable       bool      `json:"car_available" form:"CarAvailable"`                     //form
	Urgency            int       `json:"urgency" form:"Urgency"`                                //form
	RedRequired        int       `json:"red-required" form:"RedRequired"`                       //form
	RedReceived        int       `json:"red-recieved" form:"RedReceived"`
	PlateletRequired   int       `json:"platelet-required" form:"PlateletRequired"` //form
	PlateletReceived   int       `json:"platelet-recieved" form:"PlateletReceived"`
	Relationship       string    `json:"relationship" form:"Relationship"`               //form
	AcceptsRedCrescent bool      `json:"accepts_red_crescent" form:"AcceptsRedCrescent"` //form
	RedCrescentCode    string    `json:"red_crescent_code" form:"RedCrescentCode"`       //form
	DeathStatus        bool      `json:"death_status" form:"DeathStatus"`
	DeathDate          time.Time `json:"death_date,omitempty" form:"DeathDate" time_format:"2006-01-02"`
	SpecialPatient     bool      `json:"special_patient" form:"SpecialPatient"`
	Verify             string    `json:"verify,omitempty"`
}
