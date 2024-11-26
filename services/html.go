package services

import (
	"net/http"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

// Main and Dashboard pages
func Main_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main Page"})
}

func Dashboard_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"title": "Dashboard"})
}

// Patient and Donor forms
func ShowPatientForm(c *gin.Context) {
	c.HTML(http.StatusOK, "patient_form.html", gin.H{
		"title": "Patient Registration",
	})
}

func ShowDonorForm(c *gin.Context) {
	c.HTML(http.StatusOK, "donor_form.html", gin.H{
		"title": "Donor Registration",
	})
}

// Admin forms
func ShowMatchForm(c *gin.Context) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "match.html", gin.H{
		"title":    "Match Donor and Patient",
		"patients": patients,
	})
}

func ShowDonationForm(c *gin.Context) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "donation_form.html", gin.H{
		"donors":   donors,
		"patients": patients,
	})
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

// Schedule Request
func ShowSpecialPatientForm(c *gin.Context) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "scheduale_request_form.html", gin.H{
		"title":    "Special Patient Registration",
		"patients": patients,
	})
}

func ShowSchedualedRequestsPage(c *gin.Context) {
	requests, err := repositories.GetAllSchedualedRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "schedualed_requests.html", gin.H{
		"title":    "Special Requests",
		"requests": requests,
	})
}

// Updating the patient and donor
func ShowUpdatePatientForm(c *gin.Context) {
	id := c.Param("id")
	patientID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "update_patient.html", gin.H{
		"title":   "Update Patient",
		"patient": patient,
	})
}

func ShowUpdateDonorForm(c *gin.Context) {
	id := c.Param("id")
	donorID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid donor ID"})
		return
	}
	donor, err := repositories.GetDonorByID(donorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "update_donor.html", gin.H{
		"title": "Update Donor",
		"donor": donor,
	})
}

// Patients and Donors list
func ShowPatientsPage(c *gin.Context) {
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

func ShowDonorsPage(c *gin.Context) {
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

// Waiting patients
// func ShowPatientsWaitingPage(c *gin.Context) {
