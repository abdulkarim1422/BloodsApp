package services

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var unverifiedDonors = make(map[string]Donor)
var unverifiedPatients = make(map[string]Patient)

func GetDonors(c *gin.Context) {
	c.HTML(http.StatusOK, "donors.html", gin.H{
		"title":  "Donors List",
		"donors": donors,
	})
}

func CreateDonor(c *gin.Context) {
	var newDonor Donor
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
	donor.ID = len(donors) + 1
	donor.CreatedAt = time.Now()
	donor.UpdatedAt = time.Now()
	donors = append(donors, donor)

	// Remove from unverified donors
	delete(unverifiedDonors, donorID)

	c.JSON(http.StatusOK, gin.H{"message": "Verification successful"})
}

func DonorByID(c *gin.Context) {
	id := c.Param("id")
	d, err := GetDonorByID(id)
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

func GetDonorByID(id string) (*Donor, error) {
	for i := range donors {
		if strconv.Itoa(donors[i].ID) == id {
			return &donors[i], nil
		}
	}
	return nil, nil
}

func GetPatients(c *gin.Context) {
	c.HTML(http.StatusOK, "patients.html", gin.H{
		"title":    "Patients List",
		"patients": patients,
	})
}

func CreatePatient(c *gin.Context) {
	var newPatient Patient
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

	patient, exists := unverifiedPatients[patientID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	if patient.Verify != request.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Verification successful, clear the verification code and add to patients
	patient.Verify = ""
	patient.ID = len(patients) + 1
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()
	patients = append(patients, patient)

	// Remove from unverified patients
	delete(unverifiedPatients, patientID)

	c.JSON(http.StatusOK, gin.H{"message": "Verification successful"})
}

func PatientByID(c *gin.Context) {
	id := c.Param("id")
	p, err := GetPatientByID(id)
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

func GetPatientByID(id string) (*Patient, error) {
	for _, p := range patients {
		if strconv.Itoa(p.ID) == id {
			return &p, nil
		}
	}
	return nil, nil
}
