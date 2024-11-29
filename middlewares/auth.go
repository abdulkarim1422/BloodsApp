package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT token from the Authorization header in the Cookie
		tokenString, err := c.Cookie("Autherization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie required"})
			c.Abort()
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			log.Fatal(err)
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Chack the expiration time
			if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Find the user in the database
			user, err := repositories.GetUserByIdentifier(claims["username"].(string))
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Attach to req
			c.Set("user", user)

			// Continue
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ValidateRequestToken(c *gin.Context) {
	// Get the token from the query string
	tokenString := c.Param("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the expiration time
		if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Attach to req
		c.Set("token", token)

		// Continue
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
	}
}
