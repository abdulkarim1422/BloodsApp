package services

import (
	"net/http"
	"os"
	"time"

	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Login(c *gin.Context) {
	var request struct {
		Identifier string `json:"identifier"` // Can be either username or email
		Password   string `json:"password"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := initializers.DB.Where("username = ? OR email = ?", request.Identifier, request.Identifier).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token in a cookie
	// Parameters: name, value, maxAge (in seconds), path, domain, secure, httpOnly
	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	// Clear the token cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func Signup(c *gin.Context) {
	var request struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		AdminToken string `json:"admin_token"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check the signup token
	preSharedToken := os.Getenv("ADMIN_TOKEN")
	if request.AdminToken != preSharedToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin token"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user
	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUsers(c *gin.Context) {
	var request struct {
		AdminToken string `json:"admin_token"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check the admin token
	adminToken := os.Getenv("ADMIN_TOKEN")
	if request.AdminToken != adminToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin token"})
		return
	}

	var users []models.User
	if err := initializers.DB.Select("username", "email").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
func ProtectedRoute(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "username": username})
}
