package services

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

// Request --------------------------

func GetAllRequests(c *gin.Context) {
	requests, err := repositories.GetAllRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requests)
}

func MarkAsDonated(requestID int, donationType string, feedback string) error {
	request, err := repositories.GetRequestByID(requestID)
	if err != nil {
		return fmt.Errorf("request not found")
	}
	if request == nil {
		return fmt.Errorf("request not found")
	}

	// Mark the request as completed (marked by the donor)
	request.MarkedAsCompleted = true
	request.DonorFeedback = feedback

	// Mark the donation type as received
	if donationType == "red" {
		request.RedReceived = true
	} else if donationType == "platelet" {
		request.PlateletReceived = true
	}

	// Update the request
	if err := repositories.UpdateRequest(request); err != nil {
		return fmt.Errorf("failed to update request")
	}

	// send feedback to the patient
	go RecievePetientFeedbackInfo(request)

	return nil
}

func MarkAsReceived(requestID int, qqq string, feedback string) error {
	request, err := repositories.GetRequestByID(requestID)
	if err != nil {
		return fmt.Errorf("request not found")
	}
	if request == nil {
		return fmt.Errorf("request not found")
	}

	// Mark the request as accepted or rejected (marked by the patient)
	request.PatientFeedback = feedback
	if qqq == "yes" {
		request.RequestAccepted = true
	} else if qqq == "no" {
		request.RequestRejected = true
	}

	// Update the request
	if err := repositories.UpdateRequest(request); err != nil {
		return fmt.Errorf("failed to update request")
	}

	return nil
}

// SchedualedRequest --------------------------

func CreateSchedualedRequest(c *gin.Context) {
	var schedualedRequest models.SchedualedRequest
	if err := c.ShouldBind(&schedualedRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Log the request payload
	fmt.Printf("Received request: %+v\n", schedualedRequest)

	// Get the patient
	patient, err := repositories.GetPatientByID(schedualedRequest.PatientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add patient info to the schedualedRequest
	schedualedRequest.LatinName = patient.LatinName
	schedualedRequest.PhoneNumber = patient.PhoneNumber
	schedualedRequest.BloodType = patient.BloodType

	// Create the schedualed request
	if err := repositories.CreateSchedualedRequest(&schedualedRequest); err != nil {
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
	schedualedRequests, err := repositories.GetAllSchedualedRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedualedRequests)
}

func PerformSchedualedRequest(schedualedRequestID int) error {
	schedualedRequest, err := repositories.GetSchedualedRequestByID(schedualedRequestID)
	if err != nil {
		return fmt.Errorf("schedualed request not found")
	}
	if schedualedRequest == nil {
		return fmt.Errorf("schedualed request not found")
	}

	// Get the patient
	patient, err := repositories.GetPatientByID(schedualedRequest.PatientID)
	if err != nil {
		return fmt.Errorf("patient not found")
	}

	// Update patient's reqiured blood
	patient.RedRequired = schedualedRequest.RedRequired
	patient.PlateletRequired = schedualedRequest.PlateletRequired

	if err := repositories.UpdatePatient(patient); err != nil {
		return fmt.Errorf("failed to update patient")
	}

	// Update schedualed request
	schedualedRequest.RequestsSent = schedualedRequest.RequestsSent + 1
	schedualedRequest.RequestFrequency = schedualedRequest.RequestFrequency - 1
	schedualedRequest.NextRequestDate = time.Now().AddDate(0, 0, schedualedRequest.RequestInterval)
	if err := repositories.UpdateSchedualedRequest(schedualedRequest); err != nil {
		return fmt.Errorf("failed to update schedualed request")
	}

	return nil
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
