package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedPatientRoutes(router *gin.Engine) {
	router.GET("/patients/:id", services.PatientByID)
	router.POST("/patients", services.CreatePatient)
}
