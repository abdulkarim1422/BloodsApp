package handlers

import (
	"fmt"
	"net/http"

	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func SendMatchResult(c *gin.Context) {
	var response struct {
		PatientID int    `json:"patient_id"`
		Donors    []int  `json:"donors"`
		Message   string `json:"message"`
	}

	// Bind the JSON to the struct
	if err := c.ShouldBindJSON(&response); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Access the values
	patientID := response.PatientID
	donors := response.Donors

	// Log or use the values
	fmt.Printf("Patient ID: %d\n", patientID)
	fmt.Printf("Donors: %v\n", donors)

	// Send a response
	services.SendMatchResult(patientID, donors)

	c.JSON(http.StatusOK, gin.H{"message": "تمّ انتهاء عمل البرنامج بنجاح"})
}
