package services

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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

	newDonor.ID = len(donors) + 1 // Assign a new ID
	newDonor.CreatedAt = time.Now()
	newDonor.UpdatedAt = time.Now()
	donors = append(donors, newDonor)
	c.JSON(http.StatusCreated, newDonor)
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
	for _, d := range donors {
		if strconv.Itoa(d.ID) == id {
			return &d, nil
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

	newPatient.ID = len(patients) + 1 // Assign a new ID
	patients = append(patients, newPatient)
	c.JSON(http.StatusCreated, newPatient)
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
