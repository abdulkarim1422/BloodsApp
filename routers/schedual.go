package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedSchedual(router *gin.Engine) {
	router.POST("/schedual-request", handlers.CreateSchedualedRequest)
	router.GET("/schedualed-requests", handlers.GetAllSchedualedRequests)
	router.DELETE("/schedualed-request/:id", handlers.DeleteScheduledRequest)
}
