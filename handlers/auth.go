package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// LoginRequest represents the request body for login
type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// SignupRequest represents the request body for signup
type SignupRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	AdminToken string `json:"admin_token"`
}

// AdminTokenRequest represents the request body for admin token
type AdminTokenRequest struct {
	AdminToken string `json:"admin_token"`
}

// ForgotPasswordRequest represents the request body for forgot password
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body LoginRequest true "Request"
// @Success 200 {string} string "Login successful"
// @Router /login [post]
func Login(c *gin.Context) {
	services.Login(c)
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Logout successful"
// @Router /logout [post]
func Logout(c *gin.Context) {
	services.Logout(c)
}

// Signup godoc
// @Summary Signup
// @Description Signup
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body SignupRequest true "Request"
// @Success 201 {string} string "User created successfully"
// @Router /signup [post]
func Signup(c *gin.Context) {
	services.Signup(c)
}

// GetUsers godoc
// @Summary Get users
// @Description Get users
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body AdminTokenRequest true "Request"
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	services.GetUsers(c)
}

// CheckLogin godoc
// @Summary Check login
// @Description Check login
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Login successful"
// @Router /check-login [post]
func CheckLogin(c *gin.Context) {
	services.CheckLogin(c)
}

// UserForgotPassword godoc
// @Summary Forgot Password
// @Description Send reset password link
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body ForgotPasswordRequest true "Request"
// @Success 200 {string} string "Password reset email sent successfully"
// @Router /forgot-password [post]
func UserForgotPassword(c *gin.Context) {
	services.UserForgotPassword(c)
}
