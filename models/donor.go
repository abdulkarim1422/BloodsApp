package models

import (
	"time"

	"gorm.io/gorm"
)

// Donor represents a donor in the system
// @Description Donor represents a donor in the system
type Donor struct {
	gorm.Model
	FirstName      string    `json:"first_name" form:"FirstName"`                                    //none
	LastName       string    `json:"last_name,omitempty" form:"LastName"`                            //none
	LatinName      string    `json:"latinName,omitempty" form:"LatinName"`                           //FORM~
	PhoneNumber    string    `json:"phone_number" form:"PhoneNumber"`                                //FORM~
	Email          string    `json:"email" form:"Email"`                                             //none
	BloodType      string    `json:"bloodType" form:"BloodType"`                                     //FORM~
	MotherLanguage string    `json:"mother_language" form:"MotherLanguage"`                          //none
	BirthDate      time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"` //none
	Gender         string    `json:"gender" form:"Gender"`                                           //none
	Address        Address   `json:"address,omitempty" gorm:"embedded"`                              //FORM~*
	Transportation string    `json:"transportation" form:"Transportation"`                           //FORM~
	Paid           bool      `json:"paid,omitempty"`                                                 //none
	RedTimer       time.Time `json:"red_timer" time_format:"2006-01-02"`                             //auto
	PlateletTimer  time.Time `json:"platelet_timer" time_format:"2006-01-02"`                        //auto
	Score          int       `json:"score"`                                                          //auto
	Verify         string    `json:"verify,omitempty"`                                               //auto
}
