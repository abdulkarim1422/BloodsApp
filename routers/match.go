package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedMatchRoutes(router *gin.Engine) {
	router.GET("/match-red", services.MatchRedDonorPatient)
	router.GET("/match-platelet", services.MatchPlateletDonorPatient)

	router.GET("/match-red-ignore", services.MatchRedDonorPatientIgnoreBloodType)
	router.GET("/match-platelet-ignore", services.MatchPlateletDonorPatientIgnoreBloodType)

	router.POST("/process-match", services.ProcessMatchForm)
}
