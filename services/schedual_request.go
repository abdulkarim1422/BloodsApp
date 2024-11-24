package services

import (
	"net/http"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

func RegisterSchedualedRequest(c *gin.Context) {
	var newRequest models.SchedualedRequest
	if err := c.ShouldBind(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Create the schedualed request
	if err := repositories.CreateSchedualedRequest(&newRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the patient as a special patient
	patient, err := repositories.GetPatientByID(newRequest.PatientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	patient.SpecialPatient = true

	if err := repositories.UpdatePatient(patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success
	c.JSON(http.StatusOK, gin.H{"message": "Schedualed request created successfully"})
}
