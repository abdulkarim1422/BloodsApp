package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type donor struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	BloodType     string `json:"bloodType"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	RedTimer      int    `json:"redTimer"`
	PlateletTimer int    `json:"plateletTimer"`
	Score         int    `json:"score"`
}

type patient struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	BloodType        string `json:"bloodType"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	Urgency          int    `json:"urgency"`
	RedQuantity      int    `json:"red-quantity"`
	PlateletQuantity int    `json:"platelet-quantity"`
	Capability       string `json:"capability"`
}

var donors = []donor{
	{ID: 1, Name: "John Doe", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 2, Name: "Jane Doe", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 3, Name: "John Smith", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 4, Name: "Jane Smith", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 5, Name: "John Doe", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 6, Name: "Jane Doe", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 7, Name: "John Smith", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 8, Name: "Jane Smith", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 9, Name: "John Doe", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
	{ID: 10, Name: "Jane Doe", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: 0, PlateletTimer: 0, Score: 0},
}

var patients = []patient{
	{ID: 1, Name: "Jane Smith", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, Capability: "red"},
	{ID: 2, Name: "John Smith", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 2, PlateletQuantity: 0, Capability: "platelet"},
	{ID: 3, Name: "Jane Doe", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, Capability: "red"},
}

func getDonors(c *gin.Context) {
	c.JSON(http.StatusOK, donors)
}

func createDonor(c *gin.Context) {
	var newDonor donor
	if err := c.BindJSON(&newDonor); err != nil {
		return
	}
	donors = append(donors, newDonor)
	c.JSON(http.StatusCreated, newDonor)
}

func donorByID(c *gin.Context) {
	id := c.Param("id")
	d, err := getDonorByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if d == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "donor not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

func getDonorByID(id string) (*donor, error) {
	for _, d := range donors {
		if strconv.Itoa(d.ID) == id {
			return &d, nil
		}
	}
	return nil, nil
}

func getPatients(c *gin.Context) {
	c.JSON(http.StatusOK, patients)
}

func createPatient(c *gin.Context) {
	var newPatient patient
	if err := c.BindJSON(&newPatient); err != nil {
		return
	}
	patients = append(patients, newPatient)
	c.JSON(http.StatusCreated, newPatient)
}

func patientByID(c *gin.Context) {
	id := c.Param("id")
	p, err := getPatientByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if p == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func getPatientByID(id string) (*patient, error) {
	for _, p := range patients {
		if strconv.Itoa(p.ID) == id {
			return &p, nil
		}
	}
	return nil, nil
}

func updatePatientRedQuantity(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
		DonorID   int `json:"donorId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var donor *donor
	for i := range donors {
		if donors[i].ID == request.DonorID {
			donor = &donors[i]
			break
		}
	}
	if donor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	if donor.RedTimer != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Donor is not eligible for donation"})
		return
	}

	if patient.BloodType != donor.BloodType {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blood types do not match"})
		return
	}

	if patient.RedQuantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No quantity needed"})
		return
	}

	patient.RedQuantity -= 1
	donor.RedTimer += 90

	c.JSON(http.StatusOK, gin.H{
		"message":       "Patient red quantity updated and donor red timer increased",
		"patient_id":    patient.ID,
		"patient_name":  patient.Name,
		"donor_id":      donor.ID,
		"donor_name":    donor.Name,
		"remaining_red": patient.RedQuantity,
	})
}

func updatePatientPlateletQuantity(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
		DonorID   int `json:"donorId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var donor *donor
	for i := range donors {
		if donors[i].ID == request.DonorID {
			donor = &donors[i]
			break
		}
	}
	if donor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	if donor.PlateletTimer != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Donor is not eligible for donation"})
		return
	}

	warning := ""
	if patient.BloodType != donor.BloodType {
		warning = "Blood types do not match"
	}

	if patient.PlateletQuantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No quantity needed"})
		return
	}

	patient.PlateletQuantity -= 1
	donor.PlateletTimer += 120

	response := gin.H{
		"message":            "Patient platelet quantity updated and donor platelet timer increased",
		"patient_id":         patient.ID,
		"patient_name":       patient.Name,
		"donor_id":           donor.ID,
		"donor_name":         donor.Name,
		"remaining_platelet": patient.PlateletQuantity,
	}

	if warning != "" {
		response["warning"] = warning
	}

	c.JSON(http.StatusOK, response)
}

func matchRedDonorPatient(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var matchedDonor *donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].RedTimer == 0 {
			matchedDonor = &donors[i]
			break
		}
	}
	if matchedDonor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donor found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donor_id":     matchedDonor.ID,
		"donor_name":   matchedDonor.Name,
		"donor_type":   matchedDonor.BloodType,
	})
	var matchedDonors []donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].RedTimer == 0 {
			matchedDonors = append(matchedDonors, donors[i])
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donors found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donors":       matchedDonors,
	})
}

func matchRedDonorPatientIgnoreBloodType(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	compatibleBloodTypes := map[string][]string{
		"O-":  {"O-", "O+", "A-", "A+", "B-", "B+", "AB-", "AB+"},
		"O+":  {"O+", "A+", "B+", "AB+"},
		"A-":  {"A-", "A+", "AB-", "AB+"},
		"A+":  {"A+", "AB+"},
		"B-":  {"B-", "B+", "AB-", "AB+"},
		"B+":  {"B+", "AB+"},
		"AB-": {"AB-", "AB+"},
		"AB+": {"AB+"},
	}

	var matchedDonors []donor
	for i := range donors {
		if donors[i].RedTimer == 0 {
			for _, compatibleType := range compatibleBloodTypes[donors[i].BloodType] {
				if compatibleType == patient.BloodType {
					matchedDonors = append(matchedDonors, donors[i])
				}
			}
		}
	}
	if len(matchedDonors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"matched-donors": matchedDonors,
	})
}

func matchPlateletDonorPatient(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var matchedDonor *donor
	for i := range donors {
		if donors[i].BloodType == patient.BloodType && donors[i].PlateletTimer == 0 {
			matchedDonor = &donors[i]
			break
		}
	}
	if matchedDonor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donor found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donor_id":     matchedDonor.ID,
		"donor_name":   matchedDonor.Name,
	})
}

func matchPlateletDonorPatientIgnoreBloodType(c *gin.Context) {
	var request struct {
		PatientID int `json:"patientId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var patient *patient
	for i := range patients {
		if patients[i].ID == request.PatientID {
			patient = &patients[i]
			break
		}
	}
	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var matchedDonor *donor
	for i := range donors {
		if donors[i].PlateletTimer == 0 {
			matchedDonor = &donors[i]
			break
		}
	}
	if matchedDonor == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No matching donor found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Matching donor found",
		"patient_id":   patient.ID,
		"patient_name": patient.Name,
		"donor_id":     matchedDonor.ID,
		"donor_name":   matchedDonor.Name,
	})
}

func main() {
	router := gin.Default()

	router.GET("/donors", getDonors)
	router.GET("/donors/:id", donorByID)

	router.GET("/patients", getPatients)
	router.GET("/patients/:id", patientByID)

	router.POST("/donors", createDonor)
	router.POST("/patients", createPatient)

	router.POST("/donate-red", updatePatientRedQuantity)
	router.POST("/donate-platelet", updatePatientPlateletQuantity)

	router.GET("/match-red", matchRedDonorPatient)
	router.GET("/match-platelet", matchPlateletDonorPatient)

	router.GET("/match-red-ignore", matchRedDonorPatientIgnoreBloodType)
	router.GET("/match-platelet-ignore", matchPlateletDonorPatientIgnoreBloodType)

	router.Run(":8080")
}
