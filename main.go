package main

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/lib"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	lib.SetupLogOutput()

	router := gin.New()

	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")

	router.LoadHTMLGlob("templates/*.html")

	router.Use(gin.Recovery(), lib.Logger())

	router.GET("/", services.Main_Page)
	router.GET("/dashboard", services.Dashboard_Page)

	router.GET("/donor_form", services.ShowDonorForm)
	router.GET("/patient_form", services.ShowPatientForm)
	router.GET("/donation_form", services.ShowDonationForm)

	router.GET("/donors", services.GetDonors)
	router.GET("/donors/:id", services.DonorByID)

	router.GET("/patients", services.GetPatients)
	router.GET("/patients/:id", services.PatientByID)

	router.POST("/donors", services.CreateDonor)
	router.POST("/patients", services.CreatePatient)

	router.POST("/donate-red", services.DonateRed)
	router.POST("/donate-platelet", services.DonatePlatelet)

	router.GET("/match-red", services.MatchRedDonorPatient)
	router.GET("/match-platelet", services.MatchPlateletDonorPatient)

	router.GET("/match-red-ignore", services.MatchRedDonorPatientIgnoreBloodType)
	router.GET("/match-platelet-ignore", services.MatchPlateletDonorPatientIgnoreBloodType)

	router.GET("/match", services.ShowMatchForm)
	router.POST("/process-match", services.ProcessMatchForm)

	router.POST("/send", services.SendMatchResult)

	router.POST("/verify-donor/:id", services.VerifyDonor)
	router.POST("/verify-patient/:id", services.VerifyPatient)

	router.Run(":8080")
}
