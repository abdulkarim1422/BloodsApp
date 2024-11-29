package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedDonationRoutes(router *gin.RouterGroup) {
	router.POST("/donate-red", handlers.DonateRed)
	router.POST("/donate-platelet", handlers.DonatePlatelet)
}
