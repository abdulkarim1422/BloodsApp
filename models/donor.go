package models

import (
	"time"

	"gorm.io/gorm"
)

type Donor struct {
	gorm.Model
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
}
