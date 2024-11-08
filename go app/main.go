package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/donors", getDonors)
	router.GET("/donors/:id", donorByID)

	router.GET("/patients", getPatients)
	router.GET("/patients/:id", patientByID)

	router.POST("/donors", createDonor)
	router.POST("/patients", createPatient)

	router.POST("/donate-red", donateRed)
	router.POST("/donate-platelet", donatePlatelet)

	router.GET("/match-red", matchRedDonorPatient)
	router.GET("/match-platelet", matchPlateletDonorPatient)

	router.GET("/match-red-ignore", matchRedDonorPatientIgnoreBloodType)
	router.GET("/match-platelet-ignore", matchPlateletDonorPatientIgnoreBloodType)

	router.Run(":8080")
}
