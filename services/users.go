package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDonors(c *gin.Context) {
	c.JSON(http.StatusOK, donors)
}

func CreateDonor(c *gin.Context) {
	var newDonor Donor
	if err := c.BindJSON(&newDonor); err != nil {
		return
	}
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
	c.JSON(http.StatusOK, patients)
}

func CreatePatient(c *gin.Context) {
	var newPatient Patient
	if err := c.BindJSON(&newPatient); err != nil {
		return
	}
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
