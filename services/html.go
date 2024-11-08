package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main Page"})
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
