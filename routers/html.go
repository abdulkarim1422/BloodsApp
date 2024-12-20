package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterOpenHTMLRoutes(router *gin.Engine) {
	router.GET("/", services.Main_Page)
	router.GET("/donor_form", services.ShowDonorForm)
	router.GET("/patient_form", services.ShowPatientForm)
	router.GET("/login", services.ShowLoginForm)

	router.GET("/request_donor/:id", services.ShowSpecificRequestForDonor)
	router.GET("/request_patient/:id", services.ShowSpecificRequestForPatient)
}

func RegisterProtectedHTMLRoutes(router *gin.RouterGroup) {
	router.GET("/dashboard", services.Dashboard_Page)
	router.GET("/match", services.ShowMatchForm)
	router.GET("/donation_form", services.ShowDonationForm)

	router.GET("/patients", services.ShowPatientsPage)
	router.GET("/donors", services.ShowDonorsPage)

	router.GET("/scheduale_request_form", services.ShowSpecialPatientForm)
	router.GET("/schedualed_requests", services.ShowSchedualedRequestsPage)

	router.GET("/update_patient/:id", services.ShowUpdatePatientForm)
	router.GET("/update_donor/:id", services.ShowUpdateDonorForm)

	router.GET("/patients-waiting", services.ShowPatientsWaitingPage)
	router.GET("/match/:id", services.ShowOnePatientMatchForm)
}
