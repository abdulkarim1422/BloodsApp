package models

import (
	"time"

	"gorm.io/gorm"
)

// Patient represents a patient in the system
// @Description Patient represents a patient in the system
type Patient struct {
	gorm.Model
	FirstName           string    `json:"firstName" form:"FirstName"`                                     //none
	LastName            string    `json:"lastName,omitempty" form:"LastName"`                             //none
	LatinName           string    `json:"latinName,omitempty" form:"LatinName"`                           //FORM~
	PhoneNumber         string    `json:"phone_number" form:"PhoneNumber"`                                //FORM~
	Email               string    `json:"email" form:"Email"`                                             //none
	BloodType           string    `json:"bloodType" form:"BloodType"`                                     //FORM~
	BirthDate           time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"` //none
	Gender              string    `json:"gender" form:"Gender"`                                           //none
	Address             Address   `json:"address,omitempty" gorm:"embedded"`                              //FORM~*
	Transportation      string    `json:"transportation" form:"Transportation"`                           //FORM~
	Urgency             int       `json:"urgency" form:"Urgency"`                                         //FORM~
	RedRequired         int       `json:"red_required" form:"RedRequired"`                                //FORM~
	RedReceived         int       `json:"red_recieved" form:"RedReceived"`                                //auto
	PlateletRequired    int       `json:"platelet_required" form:"PlateletRequired"`                      //FORM~
	PlateletReceived    int       `json:"platelet-recieved" form:"PlateletReceived"`                      //auto
	Relationship        string    `json:"relationship" form:"Relationship"`                               //FORM~
	AcceptsRedCrescent  bool      `json:"accepts_red_crescent" form:"AcceptsRedCrescent"`                 //FORM~
	RedCrescentCode     string    `json:"red_crescent_code" form:"RedCrescentCode"`                       //FORM~
	RequiresNationality bool      `json:"requires_nationality"`                                           //FORM~
	Notes               string    `json:"notes,omitempty" form:"Notes"`                                   //none ?edit
	DeathStatus         bool      `json:"death_status" form:"DeathStatus"`                                //edit
	DeathDate           time.Time `json:"death_date,omitempty" form:"DeathDate" time_format:"2006-01-02"` //auto
	SpecialPatient      bool      `json:"special_patient" form:"SpecialPatient"`                          //Schedule form - auto
	RequestsSent        int       `json:"requests_sent" form:"RequestsSent"`                              //Match form - auto
	Verify              string    `json:"verify,omitempty"`                                               //auto
}
