package main

import (
	"github.com/abdulkarim1422/BloodsApp/middlewares"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func main() {
	middlewares.SetupLogOutput()

	router := gin.New()

	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	router.GET("/", services.Main_Page)

	router.GET("/donors", services.GetDonors)
	router.GET("/donors/:id", services.DonorByID)
	router.GET("/donor_form", services.ShowDonorForm)

	router.GET("/patients", services.GetPatients)
	router.GET("/patients/:id", services.PatientByID)
	router.GET("/patient_form", services.ShowPatientForm)

	router.POST("/donors", services.CreateDonor)
	router.POST("/patients", services.CreatePatient)

	router.POST("/donate-red", services.DonateRed)
	router.POST("/donate-platelet", services.DonatePlatelet)

	// Apply BasicAuth middleware only to match routes
	matchRoutes := router.Group("/")
	matchRoutes.Use(middlewares.BasicAuth())
	{
		matchRoutes.GET("/match-red", services.MatchRedDonorPatient)
		matchRoutes.GET("/match-platelet", services.MatchPlateletDonorPatient)
		matchRoutes.GET("/match-red-ignore", services.MatchRedDonorPatientIgnoreBloodType)
		matchRoutes.GET("/match-platelet-ignore", services.MatchPlateletDonorPatientIgnoreBloodType)
		matchRoutes.GET("/match", services.ShowMatchForm)
		matchRoutes.POST("/process-match", services.ProcessMatchForm)
	}

	router.Run(":8080")
}
