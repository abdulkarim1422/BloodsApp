package models

import "gorm.io/gorm"

// Requests represents a request in the system
// @Description Requests represents a request in the system
type Request struct {
	gorm.Model
	PatientID        int    `json:"patient_id"`
	DonorID          int    `json:"donor_id,omitempty"`
	HospitalID       int    `json:"hospital_id,omitempty"`
	HospitalName     string `json:"hospital_name" form:"HospitalName"` // Requests
	Red              int    `json:"red" form:"Red"`
	Platelet         int    `json:"platelet" form:"Platelet"`
	CarAvailable     bool   `json:"car_available" form:"CarAvailable"`
	RedRequired      int    `json:"red-required" form:"RedRequired"`           // Requests
	RedReceived      int    `json:"red-recieved" form:"RedReceived"`           // Requests
	PlateletRequired int    `json:"platelet-required" form:"PlateletRequired"` // Requests
	PlateletReceived int    `json:"platelet-recieved" form:"PlateletReceived"` // Requests
	RequestStatus    int    `json:"request_status" form:"RequestStatus"`
	Urgency          int    `json:"urgency" form:"Urgency"` // Requests
	Relationship     int    `json:"relationship" form:"Relationship"`
	//
}

type SchedualedRequest struct {
	gorm.Model
	PatientID int `json:"patient_id"`
	DonorID   int `json:"donor_id,omitempty"`
	Frequency int `json:"frequency" form:"Frequency"`
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
