package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// MatchRedDonorPatient godoc
// @Summary Match a red blood donor with a patient
// @Description Match a red blood donor with a patient based on blood type and red blood cell donation timer
// @Tags match
// @Accept json
// @Produce json
// @Param patient_id query int true "Patient ID"
// @Success 200 {object} gin.H
// @Router /match-red [get]
func MatchRedDonorPatient(c *gin.Context) {
	services.MatchRedDonorPatient(c)
}

// MatchRedDonorPatientIgnoreBloodType godoc
// @Summary Match a red blood donor with a patient (ignoring same type)
// @Description Match a red blood donor with a patient based on red blood cell donation timer
// @Tags match
// @Accept json
// @Produce json
// @Param patient_id query int true "Patient ID"
// @Success 200 {object} gin.H
// @Router /match-red-ignore [get]
func MatchRedDonorPatientIgnoreBloodType(c *gin.Context) {
	services.MatchRedDonorPatientIgnoreBloodType(c)
}

// MatchPlateletDonorPatient godoc
// @Summary Match a platelet donor with a patient
// @Description Match a platelet donor with a patient based on blood type and platelet donation timer
// @Tags match
// @Accept json
// @Produce json
// @Param patient_id query int true "Patient ID"
// @Success 200 {object} gin.H
// @Router /match-platelet [get]
func MatchPlateletDonorPatient(c *gin.Context) {
	services.MatchPlateletDonorPatient(c)
}

// MatchPlateletDonorPatientIgnoreBloodType godoc
// @Summary Match a platelet donor with a patient (ignoring same type)
// @Description Match a platelet donor with a patient based on platelet donation timer
// @Tags match
// @Accept json
// @Produce json
// @Param patient_id query int true "Patient ID"
// @Success 200 {object} gin.H
// @Router /match-platelet-ignore [get]
func MatchPlateletDonorPatientIgnoreBloodType(c *gin.Context) {
	services.MatchPlateletDonorPatientIgnoreBloodType(c)
}
