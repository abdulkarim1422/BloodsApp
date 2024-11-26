package repositories

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
	"gorm.io/gorm"
)

func GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	result := initializers.DB.Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}
	return patients, nil
}

func GetPatientByID(id int) (*models.Patient, error) {
	var patient models.Patient
	result := initializers.DB.First(&patient, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &patient, nil
	}
	return nil, nil
}

func GetPatientByNumber(number string) (*models.Patient, error) {
	var patient models.Patient
	result := initializers.DB.Where("phone_number = ?", number).First(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &patient, nil
	}
	return nil, nil
}

func CreatePatient(patient *models.Patient) error {
	result := initializers.DB.Create(patient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdatePatient(patient *models.Patient) error {
	result := initializers.DB.Model(&patient).Where("id = ?", patient.ID).Updates(models.Patient{
		FirstName: patient.FirstName,
		LastName:  patient.LastName,
		// Don't update PhoneNumber
		BloodType:        patient.BloodType,
		BirthDate:        patient.BirthDate,
		Gender:           patient.Gender,
		Address:          patient.Address,
		CarAvailable:     patient.CarAvailable,
		Urgency:          patient.Urgency,
		RedRequired:      patient.RedRequired,
		RedReceived:      patient.RedReceived,
		PlateletRequired: patient.PlateletRequired,
		PlateletReceived: patient.PlateletReceived,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeletePatient(id int) error {
	result := initializers.DB.Delete(&models.Patient{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckPatientsWaiting() ([]models.Patient, error) {
	var patients []models.Patient
	result := initializers.DB.Where("red_required != 0 OR platelet_required != 0").Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}
	return patients, nil
}

func AddOneRequestSent(patientID int) error {
	result := initializers.DB.Model(&models.Patient{}).Where("id = ?", patientID).Update("requests_sent", gorm.Expr("requests_sent + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
