package repositories

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
)

func CreateSchedualedRequest(request *models.SchedualedRequest) error {
	result := initializers.DB.Create(request)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllSchedualedRequests() ([]models.SchedualedRequest, error) {
	var requests []models.SchedualedRequest
	result := initializers.DB.Preload("Patient").Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func DeleteScheduledRequest(id int) error {
	result := initializers.DB.Delete(&models.SchedualedRequest{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
