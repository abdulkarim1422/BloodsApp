package models

import "gorm.io/gorm"

type Requests struct {
	gorm.Model
	PatientID        int    `json:"patient_id"`
	DonorID          int    `json:"donor_id,omitempty"`
	Red              int    `json:"red" form:"Red"`
	Platelet         int    `json:"platelet" form:"Platelet"`
	CarAvailable     bool   `json:"car_available" form:"CarAvailable"`
	RedRequired      int    `json:"red-required" form:"RedRequired"`           // Requests
	RedReceived      int    `json:"red-recieved" form:"RedReceived"`           // Requests
	PlateletRequired int    `json:"platelet-required" form:"PlateletRequired"` // Requests
	PlateletReceived int    `json:"platelet-recieved" form:"PlateletReceived"` // Requests
	HospitalName     string `json:"hospital_name" form:"HospitalName"`         // Requests
	RequestStatus    int    `json:"request_status" form:"RequestStatus"`
	Urgency          int    `json:"urgency" form:"Urgency"` // Requests
	Relationship     int    `json:"relationship" form:"Relationship"`
	//
}
