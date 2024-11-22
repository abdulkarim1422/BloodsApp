package routers

import (
	_ "github.com/abdulkarim1422/BloodsApp/docs" // Swagger docs | Import the generated documentation
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/abdulkarim1422/BloodsApp/lib"
	"github.com/abdulkarim1422/BloodsApp/middlewares"
	"github.com/abdulkarim1422/BloodsApp/services"

	"github.com/gin-gonic/gin"
)

// @title BloodsApp API
// @version 1
// @description This is a simple BloodsApp API server.
// @host localhost:8080
// @BasePath /swagger
func AllRouters() {
	router := gin.New()

	RegisterStaticRoutes(router)

	router.Use(gin.Recovery(), lib.Logger())

	// Public routes
	RegisterOpenHTMLRoutes(router)
	RegisterSendRoutes(router)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.POST("/check-login", services.CheckLogin)

		RegisterProtectedHTMLRoutes(router)
		RegisterProtectedPatientRoutes(router)
		RegisterProtectedDonationRoutes(router)
		RegisterProtectedMatchRoutes(router)
	}

	/* Auth API */
	RegisterAuthRoutes(router)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
