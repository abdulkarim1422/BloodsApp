package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func donateRed(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
		DonorID   int `json:"donorId"`
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

	var donor *donor
	for i := range donors {
		if donors[i].ID == request.DonorID {
			donor = &donors[i]
			break
		}
	}
	if donor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	if donor.RedTimer.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Donor is not yet eligible for donation"})
		return
	}

	if patient.BloodType != donor.BloodType {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blood types do not match"})
		return
	}

	if patient.RedQuantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No quantity needed"})
		return
	}

	patient.RedQuantity -= 1
	donor.RedTimer = time.Now().AddDate(0, 0, 90)
	donor.Score += 1

	c.JSON(http.StatusOK, gin.H{
		"message":         "Patient red quantity updated and donor red timer increased",
		"patient_id":      patient.ID,
		"patient_name":    patient.Name,
		"donor_id":        donor.ID,
		"donor_name":      donor.Name,
		"donor-score":     donor.Score,
		"donor-red-timer": donor.RedTimer,
		"remaining_red":   patient.RedQuantity,
	})
}

func donatePlatelet(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
		DonorID   int `json:"donorId"`
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

	var donor *donor
	for i := range donors {
		if donors[i].ID == request.DonorID {
			donor = &donors[i]
			break
		}
	}
	if donor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	if donor.PlateletTimer.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Donor is not yet eligible for donation"})
		return
	}

	warning := ""
	if patient.BloodType != donor.BloodType {
		warning = "Blood types do not match"
	}

	if patient.PlateletQuantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No quantity needed"})
		return
	}

	patient.PlateletQuantity -= 1
	donor.PlateletTimer = time.Now().AddDate(0, 0, 7)
	donor.Score += 1

	response := gin.H{
		"message":             "Patient platelet quantity updated and donor platelet timer increased",
		"patient_id":          patient.ID,
		"patient_name":        patient.Name,
		"donor_id":            donor.ID,
		"donor_name":          donor.Name,
		"donor-score":         donor.Score,
		"donor-platlet-timer": donor.PlateletTimer,
		"remaining_platelet":  patient.PlateletQuantity,
	}

	if warning != "" {
		response["warning"] = warning
	}

	c.JSON(http.StatusOK, response)
}
