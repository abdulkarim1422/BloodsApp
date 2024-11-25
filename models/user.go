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
	Last_activity time.Time `json:"last_activity" time_format:"2006-01-02"`
}

// Session represents a session in the system
// @Description Session represents a session in the system
type Session struct {
	gorm.Model
	Username   string    `json:"username"`
	IP_address string    `json:"ip_address"`
	LoginTime  time.Time `json:"login_time"  time_format:"2006-01-02"`
	LogoutTime time.Time `json:"logout_time" time_format:"2006-01-02"`
}
