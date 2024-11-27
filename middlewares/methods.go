package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MethodOverride() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			if override := c.PostForm("_method"); override == "DELETE" {
				c.Request.Method = "DELETE"
				fmt.Println("Overriding POST to DELETE for URL:", c.Request.URL.Path)
			}
		}
	}
}
