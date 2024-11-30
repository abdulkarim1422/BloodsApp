package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterOpenPatientRoutes(router *gin.Engine) {
	router.POST("/patient", handlers.CreatePatient)
}

func RegisterProtectedPatientRoutes(router *gin.RouterGroup) {
	router.GET("/patient/:id", handlers.GetPatientByID)
	router.PUT("/patient/:id", handlers.UpdatePatient)
	router.DELETE("/patient/:id", handlers.DeletePatient)
	router.GET("/waiting-patients", handlers.CheckPatientsWaiting)

	router.POST("/api/patient", handlers.ApiCreatePatient)
}
