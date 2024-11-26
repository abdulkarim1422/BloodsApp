package models

import (
	"time"

	"gorm.io/gorm"
)

// // Requests represents a request in the system
// // @Description Requests represents a request in the system
// type Request struct {
// 	gorm.Model
// 	PatientID    int    `json:"patient_id"`
// 	DonorID      int    `json:"donor_id,omitempty"`
// 	HospitalID   int    `json:"hospital_id,omitempty"`
// 	HospitalName string `json:"hospital_name" form:"HospitalName"` // Request
// 	Red          int    `json:"red" form:"Red"`
// 	Platelet     int    `json:"platelet" form:"Platelet"`
// 	CarAvailable bool   `json:"car_available" form:"CarAvailable"`
// 	RedRequired  int    `json:"red_required" form:"RedRequired"` // Request
// 	// RedReceived        int    `json:"red_recieved" form:"RedReceived"`           // Patient
// 	PlateletRequired int `json:"platelet_required" form:"PlateletRequired"` // Request
// 	// PlateletReceived   int    `json:"platelet_recieved" form:"PlateletReceived"` // Patient
// 	RequestStatus      int    `json:"request_status" form:"RequestStatus"`
// 	Urgency            int    `json:"urgency" form:"Urgency"` // Request
// 	Relationship       int    `json:"relationship" form:"Relationship"`
// 	AcceptsRedCrescent bool   `json:"accepts_red_crescent" form:"AcceptsRedCrescent"`
// 	RedCrescentCode    string `json:"red_crescent_code" form:"RedCrescentCode"`
// 	//
// }

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
