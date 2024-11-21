package routers

import (
	"github.com/abdulkarim1422/BloodsApp/lib"
	"github.com/abdulkarim1422/BloodsApp/middlewares"
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

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

	router.Run(":8080")
}
