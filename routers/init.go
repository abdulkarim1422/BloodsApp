package routers

import (
	"net/http"

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

	router.Use(gin.Recovery(), lib.Logger(), middlewares.MethodOverride())

	// Public routes
	RegisterOpenHTMLRoutes(router)
	RegisterSendRoutes(router)
	RegisterOpenPatientRoutes(router)
	RegisterOpenDonorRoutes(router)
	RegisterOpenScheduale(router)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.POST("/check-login", services.CheckLogin)

		RegisterProtectedHTMLRoutes(protected)
		RegisterProtectedDonationRoutes(protected)
		RegisterProtectedDonorRoutes(protected)
		RegisterProtectedPatientRoutes(protected)
		RegisterProtectedMatchRoutes(protected)
		RegisterProtectedScheduale(protected)
	}

	/* Auth API */
	RegisterAuthRoutes(router)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handle 404 errors
	router.NoRoute(func(c *gin.Context) {
		services.RenderErrorPage(c, http.StatusNotFound, "Page not found")
	})

	router.Run(":8080")
}
