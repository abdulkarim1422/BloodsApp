package routers

import (
	"github.com/abdulkarim1422/BloodsApp/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)
	router.POST("/logout", handlers.Logout)
	router.POST("/signup", handlers.Signup)
	router.GET("/users", handlers.GetAllUsers)
	router.POST("/forgot-password", handlers.UserForgotPassword)

	router.GET("/sessions", handlers.GetAllSessions)
}
