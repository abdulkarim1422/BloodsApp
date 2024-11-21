package routers

import "github.com/gin-gonic/gin"

func RegisterStaticRoutes(router *gin.Engine) {
	/* CSS & JS */
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	/* HTML */
	router.LoadHTMLGlob("templates/*.html")
}
