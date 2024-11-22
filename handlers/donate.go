package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// DonateBlood godoc
// @Summary Donate blood to a patient
// @Description Donate blood to a patient
// @ID donate-blood
// @Accept  json
// @Produce  json
// @Router /donate-red [post]
func DonateRed(c *gin.Context) {
	services.DonateRed(c)
}

// DonatePlatelet godoc
// @Summary Donate platelet to a patient
// @Description Donate platelet to a patient
// @ID donate-platelet
// @Accept  json
// @Produce  json
// @Router /donate-platelet [post]
func DonatePlatelet(c *gin.Context) {
	services.DonatePlatelet(c)
}
