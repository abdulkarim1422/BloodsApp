package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

var compatibleBloodTypes = map[string][]string{
	"O-":  {"O-", "O+", "A-", "A+", "B-", "B+", "AB-", "AB+"},
	"O+":  {"O+", "A+", "B+", "AB+"},
	"A-":  {"A-", "A+", "AB-", "AB+"},
	"A+":  {"A+", "AB+"},
	"B-":  {"B-", "B+", "AB-", "AB+"},
	"B+":  {"B+", "AB+"},
	"AB-": {"AB-", "AB+"},
	"AB+": {"AB+"},
}

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

func MatchRedDonorPatientAnyBloodType(patientID int) ([]models.Donor, error) {
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
		if donors[i].RedTimer.Before(time.Now()) {
			matchedDonors = append(matchedDonors, donors[i])
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
			for _, compatibleType := range compatibleBloodTypes[patient.BloodType] {
				if compatibleType == donors[i].BloodType {
					matchedDonors = append(matchedDonors, donors[i])
					break
				}
			}
		}
	}
	if len(matchedDonors) == 0 {
		return nil, fmt.Errorf("no matching donor found")
	}

	return matchedDonors, nil
}

func MatchPlateletDonorPatientAnyBloodType(patientID int) ([]models.Donor, error) {
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
		if request.IgnoreBloodType == "any" {
			donors, err_match = MatchRedDonorPatientAnyBloodType(request.PatientID)
		} else if request.IgnoreBloodType == "ignore" {
			donors, err_match = MatchRedDonorPatientIgnoreBloodType(request.PatientID)
		} else {
			donors, err_match = MatchRedDonorPatient(request.PatientID)
		}
	} else if request.MatchType == "platelet" {
		if request.IgnoreBloodType == "any" {
			donors, err_match = MatchPlateletDonorPatientAnyBloodType(request.PatientID)
		} else if request.IgnoreBloodType == "ignore" {
			donors, err_match = MatchPlateletDonorPatientIgnoreBloodType(request.PatientID)
		} else {
			donors, err_match = MatchPlateletDonorPatient(request.PatientID)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match type"})
		return
	}
	if err_match != nil {
		c.HTML(http.StatusOK, "match_result.html", gin.H{
			"message": fmt.Sprintf("No matching donors found, %v", err_match),
			"patient": patient,
		})
		return
	}

	c.HTML(http.StatusOK, "match_result.html", gin.H{
		"message":    "Matching donors found",
		"patient":    patient,
		"patient_id": patient.ID,
		"donors":     donors,
	})
}
