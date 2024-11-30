package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterOpenDonorRoutes(router *gin.Engine) {
	router.POST("/donor", handlers.CreateDonor)
}

func RegisterProtectedDonorRoutes(router *gin.RouterGroup) {
	router.GET("/donor/:id", handlers.DonorByID)
	router.PUT("/donor/:id", handlers.UpdateDonor)
	router.DELETE("/donor/:id", handlers.DeleteDonor)
}
