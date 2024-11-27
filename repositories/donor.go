package repositories

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
)

func GetAllDonors() ([]models.Donor, error) {
	var donors []models.Donor
	result := initializers.DB.Find(&donors)
	if result.Error != nil {
		return nil, result.Error
	}
	return donors, nil
}

func GetDonorByID(id int) (*models.Donor, error) {
	var donor models.Donor
	result := initializers.DB.First(&donor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &donor, nil
	}
	return nil, nil
}

func GetDonorByNumber(number string) (*models.Donor, error) {
	var donor models.Donor
	result := initializers.DB.Where("phone_number = ?", number).First(&donor)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &donor, nil
	}
	return nil, nil
}

func CreateDonor(donor *models.Donor) error {
	result := initializers.DB.Create(donor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateDonor(donor *models.Donor) error {
	result := initializers.DB.Model(&donor).Where("id = ?", donor.ID).Updates(models.Donor{
		FirstName: donor.FirstName,
		LastName:  donor.LastName,
		// Don't update PhoneNumber
		BloodType:      donor.BloodType,
		BirthDate:      donor.BirthDate,
		Gender:         donor.Gender,
		Address:        donor.Address,
		Transportation: donor.Transportation,
		Paid:           donor.Paid,
		RedTimer:       donor.RedTimer,
		PlateletTimer:  donor.PlateletTimer,
		Score:          donor.Score,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteDonor(id int) error {
	result := initializers.DB.Delete(&models.Donor{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
