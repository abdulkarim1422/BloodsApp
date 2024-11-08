package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func matchRedDonorPatient(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var matchedDonor *donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].RedTimer.Before(time.Now()) {
			matchedDonor = &donors[i]
			break
		}
	}
	if matchedDonor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donor found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donor_id":     matchedDonor.ID,
		"donor_name":   matchedDonor.Name,
		"donor_type":   matchedDonor.BloodType,
	})
	var matchedDonors []donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].RedTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donors found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donors":       matchedDonors,
	})
}

func matchRedDonorPatientIgnoreBloodType(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	compatibleBloodTypes := map[string][]string{
		"O-":  {"O-", "O+", "A-", "A+", "B-", "B+", "AB-", "AB+"},
		"O+":  {"O+", "A+", "B+", "AB+"},
		"A-":  {"A-", "A+", "AB-", "AB+"},
		"A+":  {"A+", "AB+"},
		"B-":  {"B-", "B+", "AB-", "AB+"},
		"B+":  {"B+", "AB+"},
		"AB-": {"AB-", "AB+"},
		"AB+": {"AB+"},
	}

	var matchedDonors []donor
	for i := range donors {
		if donors[i].RedTimer.Before(time.Now()) {
			for _, compatibleType := range compatibleBloodTypes[donors[i].BloodType] {
				if compatibleType == patient.BloodType {
					matchedDonors = append(matchedDonors, donors[i])
				}
			}
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"matched-donors": matchedDonors,
	})
}

func matchPlateletDonorPatient(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var matchedDonor *donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].PlateletTimer.Before(time.Now()) {
			matchedDonor = &donors[i]
			break
		}
	}
	if matchedDonor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donor found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donor_id":     matchedDonor.ID,
		"donor_name":   matchedDonor.Name,
	})
}

func matchPlateletDonorPatientIgnoreBloodType(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var matchedDonor *donor
	for i := range donors {
		if donors[i].PlateletTimer.Before(time.Now()) {
			matchedDonor = &donors[i]
			break
		}
	}
	if matchedDonor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donor found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donor_id":     matchedDonor.ID,
		"donor_name":   matchedDonor.Name,
	})
}
