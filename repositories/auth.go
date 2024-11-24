package repositories

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
)

// GetUserByIdentifier retrieves a user by username or email
func GetUserByIdentifier(identifier string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error
	return user, err
}

// CreateUser creates a new user in the database
func CreateUser(user models.User) error {
	return initializers.DB.Create(&user).Error
}

// GetAllUsers retrieves all users with selected fields
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := initializers.DB.Select("username", "email").Find(&users).Error
	return users, err
}

func CreateSession(session models.Session) error {
	return initializers.DB.Create(&session).Error
}
