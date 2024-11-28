package models

import (
	"time"

	"gorm.io/gorm"
)

// Requests represents a sent request in the system
// @Description Requests represents a sent request in the system
type Request struct {
	gorm.Model
	PatientID         int    `json:"patient_id"`
	DonorID           int    `json:"donor_id,omitempty"`
	PatientFeedback   string `json:"patient_feedback,omitempty"`                   // patient form
	DonorFeedback     string `json:"donor_feedback,omitempty"`                     // donor form
	RedReceived       bool   `json:"red_received" form:"RedReceived"`              // donor form
	PlateletReceived  bool   `json:"platelet_received" form:"PlateletReceived"`    // donor form
	MessageOpened     bool   `json:"message_opened" form:"MessageOpened"`          // donor link
	MarkedAsCompleted bool   `json:"marked_as_completed" form:"MarkedAsCompleted"` // donor form
	RequestRejected   bool   `json:"request_rejected" form:"RequestRejected"`      // patient form
	RequestAccepted   bool   `json:"request_accepted" form:"RequestAccepted"`      // patient form
	RequestCancelled  bool   `json:"request_cancelled" form:"RequestCancelled"`    // auto
}

// SchedualedRequest represents a schedualed request in the system
// @Description SchedualedRequest represents a schedualed request in the system
type SchedualedRequest struct {
	gorm.Model
	PatientID        int       `json:"patient_id" form:"patient_id"`
	LatinName        string    `json:"latinName,omitempty" form:"LatinName"`
	PhoneNumber      string    `json:"phone_number" form:"PhoneNumber"`
	BloodType        string    `json:"bloodType" form:"BloodType"`
	DiseaseName      string    `json:"disease_name" form:"disease_name"`
	RedRequired      int       `json:"red_required" form:"RedRequired"`
	PlateletRequired int       `json:"platelet_required" form:"PlateletRequired"`
	RequestInterval  int       `json:"request_interval" form:"request_interval"`
	RequestFrequency int       `json:"request_frequency" form:"request_frequency"`
	NextRequestDate  time.Time `json:"next_request_date" form:"next_request_date" time_format:"2006-01-02"`
	RequestsSent     int       `json:"requests_sent" form:"requests_sent"`
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
