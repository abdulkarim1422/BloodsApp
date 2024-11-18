package services

import (
	"net/http"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

var unverifiedDonors = make(map[string]models.Donor)
var unverifiedPatients = make(map[string]models.Patient)

func GetDonors(c *gin.Context) {
	var donors []models.Donor
	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "donors.html", gin.H{
		"title":  "Donors List",
		"donors": donors,
	})
}

func CreateDonor(c *gin.Context) {
	var newDonor models.Donor
	if err := c.ShouldBind(&newDonor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle the CarAvailable checkbox value
	newDonor.CarAvailable = c.PostForm("CarAvailable") == "true"

	// Generate and assign verification code
	newDonor.Verify = generateVerificationCode()

	// Send verification code via WhatsApp
	go sendVerificationCode(newDonor.PhoneNumber, newDonor.Verify)

	// Temporarily store the donor data
	tempID := strconv.Itoa(len(unverifiedDonors) + 1)
	unverifiedDonors[tempID] = newDonor

	c.JSON(http.StatusCreated, gin.H{"id": tempID})
}

func VerifyDonor(c *gin.Context) {
	donorID := c.Param("id")
	var request struct {
		VerificationCode string `json:"verificationCode"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	donor, exists := unverifiedDonors[donorID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	if donor.Verify != request.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Verification successful, clear the verification code and add to donors
	donor.Verify = ""

	// Create a new donor in the database
	create_err := repositories.CreateDonor(&donor)
	if create_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": create_err.Error()})
		return
	}

	// Remove from unverified donors after successful verification and commit
	delete(unverifiedDonors, donorID)

	c.JSON(http.StatusOK, gin.H{"message": "Verification successful"})
}

func DonorByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	d, err := repositories.GetDonorByID(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if d == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "donor not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "patients.html", gin.H{
		"title":    "Patients List",
		"patients": patients,
	})
}

func CreatePatient(c *gin.Context) {
	var newPatient models.Patient
	if err := c.ShouldBind(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Manually set `CarAvailable` based on form submission
	newPatient.CarAvailable = c.PostForm("CarAvailable") == "true"

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

func PatientByID(c *gin.Context) {
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
