package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedPatientRoutes(router *gin.RouterGroup) {
	router.GET("/patient/:id", handlers.GetPatientByID)
	router.POST("/patient", handlers.CreatePatient)
	router.PUT("/patient/:id", handlers.UpdatePatient)
	router.DELETE("/patient/:id", handlers.DeletePatient)
	router.GET("/waiting-patients", handlers.CheckPatientsWaiting)

	router.POST("/api/patient", handlers.ApiCreatePatient)
}
