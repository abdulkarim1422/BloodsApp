package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedDonationRoutes(router *gin.Engine) {
	router.POST("/donate-red", services.DonateRed)
	router.POST("/donate-platelet", services.DonatePlatelet)
}
