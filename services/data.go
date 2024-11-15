package services

import (
	"strings"
	"time"

	"github.com/abdulkarim1422/BloodsApp/db"
	"github.com/abdulkarim1422/BloodsApp/models"
)

func GetCompatibleBloodTypes() (map[string][]string, error) {
	compatibleBloodTypes := make(map[string][]string)

	dbConn := db.DbConn()
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT blood_type, compatible_types FROM blood_compatibilities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bloodType string
		var compatibleTypes string

		err := rows.Scan(&bloodType, &compatibleTypes)
		if err != nil {
			return nil, err
		}

		typesArray := strings.Split(compatibleTypes, ",")
		compatibleBloodTypes[bloodType] = typesArray
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return compatibleBloodTypes, nil
}

var donors = []models.Donor{
	{ID: 1, FirstName: "Donor One", LastName: "DDD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 2, FirstName: "Donor Two", LastName: "DDD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 3, FirstName: "Donor Three", LastName: "DDD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 4, FirstName: "Donor Four", LastName: "DDD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 5, FirstName: "Donor Five", LastName: "DDD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 6, FirstName: "Donor Six", LastName: "DDD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 7, FirstName: "Donor Seven", LastName: "DDD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 8, FirstName: "Donor Eight", LastName: "DDD", BloodType: "B-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 9, FirstName: "Donor Nine", LastName: "DDD", BloodType: "A-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 10, FirstName: "Donor Ten", LastName: "DDD", BloodType: "AB-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 11, FirstName: "Donor Eleven", LastName: "DDD", BloodType: "AB-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 12, FirstName: "Donor Twelve", LastName: "DDD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "+905511678290", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
}

var patients = []models.Patient{
	{ID: 1, FirstName: "Patient One", LastName: "PPP", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 1, RedReceived: 0, PlateletRequired: 1, PlateletReceived: 0, CarAvailable: true, HospitalName: "Hospital One"},
	{ID: 2, FirstName: "Patient Two", LastName: "PPP", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 2, RedReceived: 0, PlateletRequired: 0, PlateletReceived: 0, CarAvailable: false, HospitalName: "Hospital Two"},
	{ID: 3, FirstName: "Patient Three", LastName: "PPP", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", Urgency: 0, RedRequired: 1, RedReceived: 0, PlateletRequired: 1, PlateletReceived: 0, CarAvailable: true, HospitalName: "Hospital Three"},
}

var countriesList = []models.Countries{
	{ID: 1, CountryName: "Turkey"},
}

var citiesList = []models.Cities{
	{ID: 1, CodeNumber: "TR34", CityName: "Istanbul", CountryID: 1},
	{ID: 2, CodeNumber: "TR50", CityName: "Ankara", CountryID: 1},
	{ID: 3, CodeNumber: "TR51", CityName: "Izmir", CountryID: 1},
}

var districtList = []models.District{
	{ID: 1, DistrictName: "Kadikoy", CityID: 1},
	{ID: 2, DistrictName: "Besiktas", CityID: 1},
	{ID: 3, DistrictName: "Cankaya", CityID: 2},
	{ID: 4, DistrictName: "Bornova", CityID: 3},
}

var hospitalsList = []models.Hospitals{
	{ID: 1, HospitalName: "Acibadem", DistrictID: 1, PostalCode: 34710},
	{ID: 2, HospitalName: "Memorial", DistrictID: 2, PostalCode: 34330},
	{ID: 3, HospitalName: "Ankara Hospital", DistrictID: 3, PostalCode: 06510},
	{ID: 4, HospitalName: "Izmir Hospital", DistrictID: 4, PostalCode: 35030},
}
