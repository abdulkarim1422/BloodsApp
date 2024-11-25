package models

import (
	"time"

	"gorm.io/gorm"
)

// Donor represents a donor in the system
// @Description Donor represents a donor in the system
type Donor struct {
	gorm.Model
	FirstName      string    `json:"first_name" form:"FirstName"`
	LastName       string    `json:"last_name,omitempty" form:"LastName"`
	LatinName      string    `json:"latinName,omitempty" form:"LatinName"` //form
	PhoneNumber    string    `json:"phone_number" form:"PhoneNumber"`      //form
	Email          string    `json:"email" form:"Email"`
	BloodType      string    `json:"bloodType" form:"BloodType"` //form
	MotherLanguage string    `json:"mother_language" form:"MotherLanguage"`
	BirthDate      time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender         string    `json:"gender" form:"Gender"`
	Address        Address   `json:"address,omitempty" gorm:"embedded"`           //form*
	CarAvailable   bool      `json:"car_available,omitempty" form:"CarAvailable"` //form
	Paid           bool      `json:"paid,omitempty"`
	RedTimer       time.Time `json:"red_timer" time_format:"2006-01-02"`
	PlateletTimer  time.Time `json:"platelet_timer" time_format:"2006-01-02"`
	Score          int       `json:"score"`
	Verify         string    `json:"verify,omitempty"`
}
