package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterOpenHTMLRoutes(router *gin.Engine) {
	router.GET("/", services.Main_Page)
	router.GET("/donor_form", services.ShowDonorForm)
	router.GET("/patient_form", services.ShowPatientForm)
}

func RegisterProtectedHTMLRoutes(router *gin.Engine) {
	router.GET("/dashboard", services.Dashboard_Page)
	router.GET("/match", services.ShowMatchForm)
	router.GET("/donation_form", services.ShowDonationForm)

	router.GET("/patients", services.ShowPatientsPage)
	router.GET("/donors", services.ShowDonorsPage)
}