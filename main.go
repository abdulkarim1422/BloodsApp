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

	router.GET("/match-red", services.MatchRedDonorPatient)
	router.GET("/match-platelet", services.MatchPlateletDonorPatient)

	router.GET("/match-red-ignore", services.MatchRedDonorPatientIgnoreBloodType)
	router.GET("/match-platelet-ignore", services.MatchPlateletDonorPatientIgnoreBloodType)

	router.Run(":8080")
}
