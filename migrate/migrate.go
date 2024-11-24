package main

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Patient{})
	initializers.DB.AutoMigrate(&models.Donor{})

	initializers.DB.AutoMigrate(&models.Request{})
	initializers.DB.AutoMigrate(&models.SchedualedRequest{})
	initializers.DB.AutoMigrate(&models.Disease{})
	initializers.DB.AutoMigrate(&models.Feedback{})

	initializers.DB.AutoMigrate(&models.Country{})
	initializers.DB.AutoMigrate(&models.City{})
	initializers.DB.AutoMigrate(&models.District{})
	initializers.DB.AutoMigrate(&models.Hospital{})

	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Session{})
}
