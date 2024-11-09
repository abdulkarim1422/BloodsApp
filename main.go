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

	router.Use(gin.Recovery(), middlewares.Logger())

	router.GET("/", services.Main_Page)

	router.GET("/donor_form", services.ShowDonorForm)
	router.GET("/patient_form", services.ShowPatientForm)

	matchRoutes := router.Group("/")
	matchRoutes.Use(middlewares.BasicAuth())
	{
		router.GET("/donors", services.GetDonors)
		router.GET("/donors/:id", services.DonorByID)

		router.GET("/patients", services.GetPatients)
		router.GET("/patients/:id", services.PatientByID)

		router.POST("/donors", services.CreateDonor)
		router.POST("/patients", services.CreatePatient)

		router.POST("/donate-red", services.DonateRed)
		router.POST("/donate-platelet", services.DonatePlatelet)

		matchRoutes.GET("/match-red", services.MatchRedDonorPatient)
		matchRoutes.GET("/match-platelet", services.MatchPlateletDonorPatient)

		matchRoutes.GET("/match-red-ignore", services.MatchRedDonorPatientIgnoreBloodType)
		matchRoutes.GET("/match-platelet-ignore", services.MatchPlateletDonorPatientIgnoreBloodType)

		matchRoutes.GET("/match", services.ShowMatchForm)
		matchRoutes.POST("/process-match", services.ProcessMatchForm)
	}

	router.Run(":8080")
}
