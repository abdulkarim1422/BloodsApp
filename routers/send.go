package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterSendRoutes(router *gin.Engine) {
	router.POST("/send", services.SendMatchResult)

	router.POST("/verify-donor/:id", services.VerifyDonor)
	router.POST("/verify-patient/:id", services.VerifyPatient)
}
