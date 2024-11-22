package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedPatientRoutes(router *gin.Engine) {
	router.GET("/patients/:id", handlers.PatientByID)
	router.POST("/patients", handlers.CreatePatient)
}
