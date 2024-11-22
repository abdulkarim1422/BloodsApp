package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedMatchRoutes(router *gin.Engine) {
	router.GET("/match-red", handlers.MatchRedDonorPatient)
	router.GET("/match-platelet", handlers.MatchPlateletDonorPatient)

	router.GET("/match-red-ignore", handlers.MatchRedDonorPatientIgnoreBloodType)
	router.GET("/match-platelet-ignore", handlers.MatchPlateletDonorPatientIgnoreBloodType)

	router.POST("/process-match", services.ProcessMatchForm)
}
