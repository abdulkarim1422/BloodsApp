package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedMatchRoutes(router *gin.RouterGroup) {
	router.POST("/process-match", services.ProcessMatchForm)
}
