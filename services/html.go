package services

import (
	"net/http"

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
	c.HTML(http.StatusOK, "match.html", gin.H{
		"title":    "Match Donor and Patient",
		"patients": patients,
	})
}

func ShowDonationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "donation_form.html", gin.H{
		"donors":   donors,   // Assuming donors is a slice of donor objects
		"patients": patients, // Assuming patients is a slice of patient objects
	})
}
