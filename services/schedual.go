package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

func CreateSchedualedRequest(c *gin.Context) {
	var newRequest models.SchedualedRequest
	if err := c.ShouldBind(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Log the request payload
	fmt.Printf("Received request: %+v\n", newRequest)

	// Get the patient
	patient, err := repositories.GetPatientByID(newRequest.PatientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Bind the patient to the newRequest.Patient
	newRequest.Patient = *patient

	// Create the schedualed request
	if err := repositories.CreateSchedualedRequest(&newRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the patient as a special patient
	patient.SpecialPatient = true

	if err := repositories.UpdatePatient(patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success
	c.JSON(http.StatusOK, gin.H{"message": "Schedualed request created successfully"})
}

func GetAllSchedualedRequests(c *gin.Context) {
	requests, err := repositories.GetAllSchedualedRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requests)
}

func DeleteScheduledRequest(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = repositories.DeleteScheduledRequest(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Scheduled request deleted successfully"})
}
