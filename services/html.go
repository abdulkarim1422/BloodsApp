package services

import (
	"net/http"

	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

func Main_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main Page"})
}

func Dashboard_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"title": "Dashboard"})
}

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
