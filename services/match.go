package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

func MatchRedDonorPatient(patientID int) ([]models.Donor, error) {
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, fmt.Errorf("patient not found")
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		return nil, err
	}

	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].RedTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}

	if len(matchedDonors) == 0 {
		return nil, fmt.Errorf("no matching donors found")
	}

	return matchedDonors, nil
}

func MatchRedDonorPatientIgnoreBloodType(patientID int) ([]models.Donor, error) {
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, fmt.Errorf("patient not found")
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
		return nil, err
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
		return nil, fmt.Errorf("no matching donors found")
	}

	return matchedDonors, nil
}

func MatchPlateletDonorPatient(patientID int) ([]models.Donor, error) {
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, fmt.Errorf("patient not found")
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		return nil, err
	}

	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].PlateletTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}
	if len(matchedDonors) == 0 {
		return nil, fmt.Errorf("no matching donor found")
	}

	return matchedDonors, nil
}

func MatchPlateletDonorPatientIgnoreBloodType(patientID int) ([]models.Donor, error) {
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, fmt.Errorf("patient not found")
	}

	donors, err := repositories.GetAllDonors()
	if err != nil {
		return nil, err
	}

	var matchedDonors []models.Donor
	for i := range donors {
		if donors[i].PlateletTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}
	if len(matchedDonors) == 0 {
		return nil, fmt.Errorf("no matching donor found")
	}

	return matchedDonors, nil
}

func OldProcessMatchForm(c *gin.Context) {
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

	patient, err := repositories.GetPatientByID(request.PatientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var donors []models.Donor
	var err_match error
	if request.MatchType == "red" {
		if request.IgnoreBloodType == "yes" {
			donors, err_match = MatchRedDonorPatientIgnoreBloodType(request.PatientID)
		} else {
			donors, err_match = MatchRedDonorPatient(request.PatientID)
		}
	} else if request.MatchType == "platelet" {
		if request.IgnoreBloodType == "yes" {
			donors, err_match = MatchPlateletDonorPatientIgnoreBloodType(request.PatientID)
		} else {
			donors, err_match = MatchPlateletDonorPatient(request.PatientID)
		}
	}
	if err_match != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "match_result.html", gin.H{
		"message": "Matching donors found",
		"patient": patient,
		"donors":  donors,
	})
}
