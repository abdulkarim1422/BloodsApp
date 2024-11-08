package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BasicAuth() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")

	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}
