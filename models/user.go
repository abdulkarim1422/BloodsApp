package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
// @Description User represents a user in the system
type User struct {
	gorm.Model
	Username      string    `gorm:"unique;not null"`
	Email         string    `gorm:"unique;not null"`
	Password      string    `gorm:"not null"`
	Last_activity time.Time `json:"last_activity"`
}

// Session represents a session in the system
// @Description Session represents a session in the system
type Session struct {
	gorm.Model
	UserID     int       `json:"user_id"`
	IP_address string    `json:"ip_address"`
	LoginTime  time.Time `json:"login_time"`
	LogoutTime time.Time `json:"logout_time"`
}
