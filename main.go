package main

import (
	"github.com/abdulkarim1422/BloodsApp/db"
	"github.com/abdulkarim1422/BloodsApp/lib"
	"github.com/abdulkarim1422/BloodsApp/middlewares"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func main() {
	lib.SetupLogOutput()

	dbConn := db.DbConn()
	defer dbConn.Close()

	router := gin.New()

	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")

	router.LoadHTMLGlob("templates/*.html")

	router.Use(gin.Recovery(), lib.Logger())

	router.GET("/", services.Main_Page)

	router.GET("/donor_form", services.ShowDonorForm)
	router.GET("/patient_form", services.ShowPatientForm)
	router.GET("/donation_form", services.ShowDonationForm)

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

		matchRoutes.POST("/send", services.SendMatchResult)

		router.POST("/verify-donor/:id", services.VerifyDonor)
		router.POST("/verify-patient/:id", services.VerifyPatient)
	}

	router.Run(":8080")
}
