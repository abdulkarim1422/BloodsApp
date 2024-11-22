package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// DonateRed godoc
// @Summary Donate Red Blood to a patient
// @Description Donate Red Blood to a patient
// @Tags donate
// @ID donate-red
// @Accept  json
// @Produce  json
// @Param patient_id query int true "Patient ID"
// @Param donor_id query int true "Donor ID"
// @Success 200 {object} gin.H
// @Router /donate-red [post]
func DonateRed(c *gin.Context) {
	services.DonateRed(c)
}

// DonatePlatelet godoc
// @Summary Donate blood to a patient
// @Description Donate blood to a patient
// @Tags donate
// @ID donate-platelet
// @Accept  json
// @Produce  json
// @Param patient_id query int true "Patient ID"
// @Success 200 {object} gin.H
// @Router /donate-platelet [post]
func DonatePlatelet(c *gin.Context) {
	services.DonatePlatelet(c)
}
