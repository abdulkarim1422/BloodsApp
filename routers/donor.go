package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterDonorOpenRoutes(router *gin.Engine) {
	router.GET("/donors/:id", services.DonorByID)
	router.POST("/donors", services.CreateDonor)
}
