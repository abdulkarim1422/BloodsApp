package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MatchRedDonorPatient(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *Patient
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

	var matchedDonor *Donor
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
		"message":            "Matching donor found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"donor_id":           matchedDonor.ID,
		"donor_first_name":   matchedDonor.FirstName,
		"donor_last_name":    matchedDonor.LastName,
		"donor_type":         matchedDonor.BloodType,
	})
	var matchedDonors []Donor
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
		"message":            "Matching donors found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"donors":             matchedDonors,
	})
}

func MatchRedDonorPatientIgnoreBloodType(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *Patient
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

	var matchedDonors []Donor
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

func MatchPlateletDonorPatient(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *Patient
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

	var matchedDonor *Donor
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
		"message":            "Matching donor found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"donor_id":           matchedDonor.ID,
		"donor_first_name":   matchedDonor.FirstName,
		"donor_last_name":    matchedDonor.LastName,
	})
}

func MatchPlateletDonorPatientIgnoreBloodType(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *Patient
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

	var matchedDonor *Donor
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
		"message":            "Matching donor found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"donor_id":           matchedDonor.ID,
		"donor_first_name":   matchedDonor.FirstName,
		"donor_last_name":    matchedDonor.LastName,
	})
}
