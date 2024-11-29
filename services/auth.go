package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Generating JWT ------------------------------------------------
func GenerateRequestJWT(requestID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour * 30) // 30 days
	claims := &jwt.MapClaims{
		"request_id": requestID,
		"ExpiresAt":  jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "failed to create token", err
	}

	return tokenString, nil
}

func GenerateUserJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"username":  username,
		"ExpiresAt": jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "failed to create token", err
	}

	return tokenString, nil
}

// Login and Logout ------------------------------------------------
func Login(c *gin.Context) {
	var request struct {
		Identifier string `json:"identifier"` // Can be either username or email
		Password   string `json:"password"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repositories.GetUserByIdentifier(request.Identifier)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := GenerateUserJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Save the session
	CreateSession(user.Username, c.ClientIP())

	// Set the token in a cookie
	// Parameters: name, value, maxAge (in seconds), path, domain, secure, httpOnly
	c.SetCookie("Autherization", token, 3600*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	// Clear the token cookie
	c.SetCookie("Autherization", "", -1, "/", "", false, true)
	// c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})

	// Redirect to the main page
	c.Redirect(http.StatusFound, "/")
}

// Signup ------------------------------------------------
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

	if err := repositories.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Admin stuff ------------------------------------------------
func GetAllUsers(c *gin.Context) {
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

	users, err := repositories.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CheckLogin(c *gin.Context) {
	token, _ := c.Cookie("Autherization")
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Logged in!", "token": token, "username": username})
}

// Forgot Password ------------------------------------------------
func UserForgotPassword(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repositories.GetUserByIdentifier(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate reset token
	resetToken, err := GenerateUserJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate reset token"})
		return
	}

	// Create reset password link
	resetLink := fmt.Sprintf("http://localhost:8099/reset-password?token=%s", resetToken)

	// Prepare email request
	emailReq := EmailRequest{
		To:      user.Email,
		Subject: "Password Reset Request",
		Body:    fmt.Sprintf("Click the link to reset your password: %s", resetLink),
	}

	// Send email
	if err := SendMail(emailReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}

// func UserResetPassword(c *gin.Context) {

// Session ------------------------------------------------
func CreateSession(username string, ClientIP string) {
	session := models.Session{
		Username:   username,
		IP_address: ClientIP,
		LoginTime:  time.Now(),
	}

	if err := repositories.CreateSession(session); err != nil {
		fmt.Println("Failed to save session")
	}
}

func GetAllSessions(c *gin.Context) {
	sessions, err := repositories.GetAllSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sessions"})
		return
	}

	c.JSON(http.StatusOK, sessions)
}
