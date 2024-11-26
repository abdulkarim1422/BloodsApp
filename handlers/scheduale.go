package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func CreateSchedualedRequest(c *gin.Context) {
	services.CreateSchedualedRequest(c)
}

func GetAllSchedualedRequests(c *gin.Context) {
	services.GetAllSchedualedRequests(c)
}

func DeleteScheduledRequest(c *gin.Context) {
	services.DeleteScheduledRequest(c)
}
