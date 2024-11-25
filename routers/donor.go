package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedDonorRoutes(router *gin.Engine) {
	router.GET("/donor/:id", handlers.DonorByID)
	router.POST("/donor", handlers.CreateDonor)
	router.PUT("/donor/:id", handlers.UpdateDonor)
	router.DELETE("/donor/:id", handlers.DeleteDonor)
}
