package repositories

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
)

// Request ---------------------------

func CreateRequest(request *models.Request) (int, error) {
	result := initializers.DB.Create(request)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(request.ID), nil
}

func GetAllRequests() ([]models.Request, error) {
	var requests []models.Request
	result := initializers.DB.Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}

func GetRequestByID(id int) (*models.Request, error) {
	var request models.Request
	result := initializers.DB.First(&request, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &request, nil
	}
	return nil, nil
}

func UpdateRequest(request *models.Request) error {
	result := initializers.DB.Model(&request).Where("id = ?", request.ID).Updates(request)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SetMessageOpenedByDonor(requestID int) error {
	result := initializers.DB.Model(&models.Request{}).Where("id = ?", requestID).Update("message_opened_by_donor", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SetMessageOpenedByPatient(requestID int) error {
	result := initializers.DB.Model(&models.Request{}).Where("id = ?", requestID).Update("message_opened_by_patient", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRequest(id int) error {
	result := initializers.DB.Delete(&models.Request{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SchedualedRequest ---------------------------

func CreateSchedualedRequest(schedualedRequest *models.SchedualedRequest) error {
	result := initializers.DB.Create(schedualedRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllSchedualedRequests() ([]models.SchedualedRequest, error) {
	var schedualedRequests []models.SchedualedRequest
	result := initializers.DB.Find(&schedualedRequests)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedualedRequests, nil
}

func DeleteScheduledRequest(id int) error {
	result := initializers.DB.Delete(&models.SchedualedRequest{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetWaitingScheduals() ([]models.SchedualedRequest, error) {
	var schedualedRequests []models.SchedualedRequest
	result := initializers.DB.Where("next_request_date < ?", "now()").Find(&schedualedRequests)
	if result.Error != nil {
		return nil, result.Error
	}
	return schedualedRequests, nil
}

func GetSchedualedRequestByID(id int) (*models.SchedualedRequest, error) {
	var schedualedRequest models.SchedualedRequest
	result := initializers.DB.First(&schedualedRequest, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &schedualedRequest, nil
	}
	return nil, nil
}

func UpdateSchedualedRequest(schedualedRequest *models.SchedualedRequest) error {
	result := initializers.DB.Model(&schedualedRequest).Where("id = ?", schedualedRequest.ID).Updates(schedualedRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteSchedualedRequest(id int) error {
	result := initializers.DB.Delete(&models.SchedualedRequest{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteAllSchedualedRequestsForPatient(patientID int) error {
	result := initializers.DB.Where("patient_id = ?", patientID).Delete(&models.SchedualedRequest{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func MarkAsCancelledAllPatientRequests(patientID int) error {
	result := initializers.DB.Model(&models.Request{}).Where(
		"patient_id = ? && marked_as_completed != true && request_rejected != true && request_accepted != true",
		patientID,
	).Update("request_cancelled", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
