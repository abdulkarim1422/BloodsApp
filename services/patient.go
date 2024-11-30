package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

var unverifiedPatients = make(map[string]models.Patient)

func CreatePatient(c *gin.Context) {
	var newPatient models.Patient
	if err := c.ShouldBind(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate and assign verification code
	newPatient.Verify = generateVerificationCode()

	// Send verification code via WhatsApp
	go sendVerificationCode(newPatient.PhoneNumber, newPatient.Verify)

	// Temporarily store the patient data
	tempID := strconv.Itoa(len(unverifiedPatients) + 1)
	unverifiedPatients[tempID] = newPatient

	// ADD REQUEST

	c.JSON(http.StatusCreated, gin.H{"id": tempID})
}

func ApiCreatePatient(c *gin.Context) {
	var newPatient models.Patient
	if err := c.ShouldBind(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repositories.CreatePatient(&newPatient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newPatient.ID})
}

func VerifyPatient(c *gin.Context) {
	patientID := c.Param("id")
	var request struct {
		VerificationCode string `json:"verificationCode"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Retrieve the unverified patient data
	patient, exists := unverifiedPatients[patientID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Check if the verification code matches
	if patient.Verify != request.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Verification successful, clear the verification code
	patient.Verify = ""

	// Create a new patient in the database
	create_err := repositories.CreatePatient(&patient)
	if create_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": create_err.Error()})
		return
	}

	// Remove from unverified patients after successful verification and commit
	delete(unverifiedPatients, patientID)

	c.JSON(http.StatusOK, gin.H{"message": "Verification successful"})
}

func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	p, err := repositories.GetPatientByID(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if p == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var updatedPatient models.Patient
	if err := c.ShouldBind(&updatedPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPatient.ID = uint(intID)
	err = repositories.UpdatePatient(&updatedPatient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Patient updated successfully")
	c.Redirect(http.StatusFound, "/patients")
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = repositories.DeletePatient(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = repositories.DeleteAllSchedualedRequestsForPatient(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully, as well as its scheduled requests"})
}

// Check if patient is waiting for donors
func CheckPatientsWaiting(c *gin.Context) {
	repositories.CheckPatientsWaiting()
}
