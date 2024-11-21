package routers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/login", services.Login)
	router.POST("/logout", services.Logout)
	router.POST("/signup", services.Signup)
	router.GET("/users", services.GetUsers)
	router.POST("/forgot-password", services.UserForgotPassword)
}
