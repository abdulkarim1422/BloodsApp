package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedScheduale(router *gin.Engine) {
	router.POST("/scheduale-request", handlers.CreateSchedualedRequest)
	router.GET("/schedualed-requests", handlers.GetAllSchedualedRequests)
	router.DELETE("/schedualed-request/:id", handlers.DeleteScheduledRequest)

	router.GET("/requests", services.GetAllRequests)
}
