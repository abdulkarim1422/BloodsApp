package models

import (
	"time"

	"gorm.io/gorm"
)

// Requests represents a sent request in the system
// @Description Requests represents a sent request in the system
type Request struct {
	gorm.Model
	PatientID         int     `json:"patient_id"`
	DonorID           int     `json:"donor_id,omitempty"`
	Address           Address `json:"address,omitempty" gorm:"embedded"`
	RedReceived       bool    `json:"red_received" form:"RedReceived"`
	PlateletReceived  bool    `json:"platelet_received" form:"PlateletReceived"`
	RedCrescentCode   string  `json:"red_crescent_code" form:"RedCrescentCode"`
	MessageOpened     bool    `json:"message_opened" form:"MessageOpened"`
	MarkedAsCompleted bool    `json:"marked_as_completed" form:"MarkedAsCompleted"`
	RequestRejected   bool    `json:"request_rejected" form:"RequestRejected"`
	RequestAccepted   bool    `json:"request_accepted" form:"RequestAccepted"`
	RequestCancelled  bool    `json:"request_cancelled" form:"RequestCancelled"`
}

// SchedualedRequest represents a schedualed request in the system
// @Description SchedualedRequest represents a schedualed request in the system
type SchedualedRequest struct {
	gorm.Model
	PatientID        int       `json:"patient_id" form:"patient_id"`
	Patient          Patient   `gorm:"foreignKey:PatientID"`
	RequestInterval  int       `json:"request_interval" form:"request_interval"`
	RequestFrequency int       `json:"request_frequency" form:"request_frequency"`
	DiseaseName      string    `json:"disease_name" form:"disease_name"`
	NextRequestDate  time.Time `json:"next_request_date" form:"next_request_date" time_format:"2006-01-02"`
	RequestsDone     int       `json:"requests_done" form:"requests_done"`
}

// Disease represents a disease in the system
// @Description Disease represents a disease in the system
type Disease struct {
	gorm.Model
	DiseaseName              string `json:"disease_name" form:"DiseaseName"`
	DiseaseType              string `json:"disease_type" form:"DiseaseType"`
	RequiresRepeatedDonation bool   `json:"requires_repeated_donation" form:"RequiresRepeatedDonation"`
}

// Feedback represents a feedback in the system
// @Description FeedBack represents a feedback in the system
type Feedback struct {
	gorm.Model
	DonorID   int    `json:"donor_id"`
	PatientID int    `json:"patient_id"`
	Rating    int    `json:"rating" form:"Rating"`
	Feedback  string `json:"feedback" form:"Feedback"`
}
