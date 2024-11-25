package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

func MatchRedDonorPatient(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Query("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}

	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].RedTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}

	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":            "Matching donors found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"patient_blood_type": patient.BloodType,
		"address":            patient.Address,
		"donors":             matchedDonors,
	})
}

func MatchRedDonorPatientIgnoreBloodType(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Query("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}

	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	compatibleBloodTypes := map[string][]string{
		"O-":  {"O-", "O+", "A-", "A+", "B-", "B+", "AB-", "AB+"},
		"O+":  {"O+", "A+", "B+", "AB+"},
		"A-":  {"A-", "A+", "AB-", "AB+"},
		"A+":  {"A+", "AB+"},
		"B-":  {"B-", "B+", "AB-", "AB+"},
		"B+":  {"B+", "AB+"},
		"AB-": {"AB-", "AB+"},
		"AB+": {"AB+"},
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].RedTimer.Before(time.Now()) {
			for _, compatibleType := range compatibleBloodTypes[patient.BloodType] {
				if compatibleType == donors[i].BloodType {
					matchedDonors = append(matchedDonors, donors[i])
					break
				}
			}
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":            "Matching donors found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"patient_blood_type": patient.BloodType,
		"address":            patient.Address,
		"donors":             matchedDonors,
	})
}

func MatchPlateletDonorPatient(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Query("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}

	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].PlateletTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":            "Matching donors found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"patient_blood_type": patient.BloodType,
		"address":            patient.Address,
		"donors":             matchedDonors,
	})
}

func MatchPlateletDonorPatientIgnoreBloodType(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Query("patient_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing patient_id query parameter"})
		return
	}

	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].PlateletTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":            "Matching donors found",
		"patient_id":         patient.ID,
		"patient_first_name": patient.FirstName,
		"patient_last_name":  patient.LastName,
		"patient_blood_type": patient.BloodType,
		"address":            patient.Address,
		"donors":             matchedDonors,
	})
}

func ProcessMatchForm(c *gin.Context) {
	var request struct {
		PatientID       int    `form:"patient_id"`
		MatchType       string `form:"match_type"`
		IgnoreBloodType string `form:"ignore_blood_type"`
	}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var endpoint string
	if request.MatchType == "red" {
		if request.IgnoreBloodType == "yes" {
			endpoint = "/match-red-ignore"
		} else {
			endpoint = "/match-red"
		}
	} else if request.MatchType == "platelet" {
		if request.IgnoreBloodType == "yes" {
			endpoint = "/match-platelet-ignore"
		} else {
			endpoint = "/match-platelet"
		}
	}

	// Create the request URL with query parameters
	url := fmt.Sprintf("http://localhost:8080%s?patient_id=%d", endpoint, request.PatientID)

	// Log the request details
	fmt.Printf("Sending request to endpoint: %s\n", url)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Add basic authentication header
	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	req.SetBasicAuth(username, password)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// Log the response details
	fmt.Printf("Received response with status: %d and body: %s\n", resp.StatusCode, string(body))

	// Check if the response status is not OK
	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]string
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			// Log the error response for debugging
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse error response", "response": string(body)})
			return
		}
		c.JSON(resp.StatusCode, errorResponse)
		return
	}

	// Parse the JSON response
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	// Return the JSON response
	c.HTML(http.StatusOK, "match_result.html", jsonResponse)
}
