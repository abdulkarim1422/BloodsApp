package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedScheduale(router *gin.RouterGroup) {
	// SchedualedRequest --------------------------
	router.POST("/scheduale-request", handlers.CreateSchedualedRequest)
	router.GET("/get-all-schedualed-requests", handlers.GetAllSchedualedRequests)
	router.DELETE("/schedualed-request/:id", handlers.DeleteScheduledRequest)
	router.POST("/perform-schedualed-request/:id", handlers.PerformSchedualedRequest)

	// Request --------------------------
	router.GET("/requests", services.GetAllRequests)

	// request page
	router.POST("/mark-donated", handlers.MarkAsDonated)
}
