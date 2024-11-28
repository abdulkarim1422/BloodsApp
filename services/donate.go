package services

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

func DonateRed(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Query("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}
	donorID, err := strconv.Atoi(c.Query("donor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}

	// Log the parsed request data
	fmt.Printf("Received DonateRed request >> patient: %+v donor: %v\n", patientID, donorID)

	// Get the patient and donor from the request
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	donor, err := repositories.GetDonorByID(donorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if donor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "donor not found"})
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
	} else if patient.RedRequired == 1 {
		patient.RedRequired = 0
	} else {
		patient.RedRequired -= 1
	}

	patient.RedReceived += 1
	donor.RedTimer = time.Now().AddDate(0, 0, 90)
	donor.Score += 1

	// Update the patient and donor in the database
	if err := repositories.UpdatePatient(patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateDonor(donor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // SEND FEEDBACK FORMS TO THE PATIENT AND DONOR
	// go SendFeedBackToPatient(patient.Email, patient.PhoneNumber)
	// go SendFeedBackToDonor(donor.Email, donor.PhoneNumber)

	c.JSON(http.StatusOK, gin.H{
		"message":            "Patient red quantity updated and donor red timer increased, donor score increased, and feedbacks sent",
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
	patientID, err := strconv.Atoi(c.Query("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}
	donorID, err := strconv.Atoi(c.Query("donor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing donorId query parameter"})
		return
	}

	// Log the parsed request data
	fmt.Printf("Received DonatePlatelet request >> patient: %+v donor: %v\n", patientID, donorID)

	// Get the patient and donor from the request
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	donor, err := repositories.GetDonorByID(donorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if donor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "donor not found"})
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
	} else if patient.PlateletRequired == 1 {
		patient.PlateletRequired = 0
	} else {
		patient.PlateletRequired -= 1
	}

	patient.PlateletReceived += 1
	donor.PlateletTimer = time.Now().AddDate(0, 0, 7)
	donor.Score += 1

	// Update the patient and donor in the database
	if err := repositories.UpdatePatient(patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateDonor(donor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // SEND FEEDBACK FORMS TO THE PATIENT AND DONOR
	// go SendFeedBackToPatient(patient.Email, patient.PhoneNumber)
	// go SendFeedBackToDonor(donor.Email, donor.PhoneNumber)

	response := gin.H{
		"message":             "Patient platelet quantity updated and donor platelet timer increased, donor score increased, and feedbacks sent",
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
