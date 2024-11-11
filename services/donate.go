package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DonateRed(c *gin.Context) {
	var request struct {
		PatientID int `json:"patient_id"`
		DonorID   int `json:"donor_id"`
	}
	if err := c.BindJSON(&request); err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Log the parsed request data
	fmt.Printf("Received DonateRed request: %+v\n", request)

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

	var donor *Donor
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

	if patient.RedRequired < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No quantity needed"})
		return
	}

	patient.RedRequired -= 1
	patient.RedReceived += 1
	donor.RedTimer = time.Now().AddDate(0, 0, 90)
	donor.Score += 1

	// SEND FEEDBACK FORMS TO THE PATIENT ANF DONOR
	// ???

	c.JSON(http.StatusOK, gin.H{
		"message":            "Patient red quantity updated and donor red timer increased",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"donor_id":           donor.ID,
		"donor_first_name":   donor.FirstName,
		"donor_last_name":    donor.LastName,
		"donor_score":        donor.Score,
		"donor_red_timer":    donor.RedTimer,
		"remaining_red":      patient.RedRequired,
	})
}

func DonatePlatelet(c *gin.Context) {
	var request struct {
		PatientID int `json:"patient_id"`
		DonorID   int `json:"donor_id"`
	}
	if err := c.BindJSON(&request); err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Log the parsed request data
	fmt.Printf("Received DonatePlatelet request: %+v\n", request)

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

	var donor *Donor
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

	if patient.PlateletRequired < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No quantity needed"})
		return
	}

	patient.PlateletRequired -= 1
	patient.PlateletReceived += 1
	donor.PlateletTimer = time.Now().AddDate(0, 0, 7)
	donor.Score += 1

	response := gin.H{
		"message":             "Patient platelet quantity updated and donor platelet timer increased",
		"patient_id":          patient.ID,
		"patient_first_name":  patient.FirstName,
		"patient_last_name":   patient.LastName,
		"donor_id":            donor.ID,
		"donor_first_name":    donor.FirstName,
		"donor_last_name":     donor.LastName,
		"donor_score":         donor.Score,
		"donor_platlet-timer": donor.PlateletTimer,
		"remaining_platelet":  patient.PlateletRequired,
	}

	if warning != "" {
		response["warning"] = warning
	}

	c.JSON(http.StatusOK, response)
}
