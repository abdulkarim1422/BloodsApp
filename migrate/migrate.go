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
	initializers.DB.AutoMigrate(&models.Requests{})

	initializers.DB.AutoMigrate(&models.Countries{})
	initializers.DB.AutoMigrate(&models.Cities{})
	initializers.DB.AutoMigrate(&models.District{})
	initializers.DB.AutoMigrate(&models.Hospitals{})

	initializers.DB.AutoMigrate(&models.User{})
}
