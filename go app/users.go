package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getDonors(c *gin.Context) {
	c.JSON(http.StatusOK, donors)
}

func createDonor(c *gin.Context) {
	var newDonor donor
	if err := c.BindJSON(&newDonor); err != nil {
		return
	}
	donors = append(donors, newDonor)
	c.JSON(http.StatusCreated, newDonor)
}

func donorByID(c *gin.Context) {
	id := c.Param("id")
	d, err := getDonorByID(id)
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

func getDonorByID(id string) (*donor, error) {
	for _, d := range donors {
		if strconv.Itoa(d.ID) == id {
			return &d, nil
		}
	}
	return nil, nil
}

func getPatients(c *gin.Context) {
	c.JSON(http.StatusOK, patients)
}

func createPatient(c *gin.Context) {
	var newPatient patient
	if err := c.BindJSON(&newPatient); err != nil {
		return
	}
	patients = append(patients, newPatient)
	c.JSON(http.StatusCreated, newPatient)
}

func patientByID(c *gin.Context) {
	id := c.Param("id")
	p, err := getPatientByID(id)
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

func getPatientByID(id string) (*patient, error) {
	for _, p := range patients {
		if strconv.Itoa(p.ID) == id {
			return &p, nil
		}
	}
	return nil, nil
}
