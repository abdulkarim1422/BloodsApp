package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// PatientByID godoc
// @Summary Get a patient by ID
// @Description Get a patient by ID
// @Tags patients
// @ID get-patient-by-id
// @Accept json
// @Produce json
// @Param id path string true "Patient ID"
// @Router /patient/{id} [get]
func GetPatientByID(c *gin.Context) {
	services.GetPatientByID(c)
}

// CreatePatient godoc
// @Summary Create a patient
// @Description Create a patient
// @Tags patients
// @ID create-patient
// @Accept json
// @Produce json
// @Param patient body models.Patient true "Patient object"
// @Router /patient [post]
func CreatePatient(c *gin.Context) {
	services.CreatePatient(c)
}

// UpdatePatient godoc
// @Summary Update a patient
// @Description Update a patient
// @Tags patients
// @ID update-patient
// @Accept json
// @Produce json
// @Param id path string true "Patient ID"
// @Param patient body models.Patient true "Patient object"
// @Router /patient/{id} [put]
func UpdatePatient(c *gin.Context) {
	services.UpdatePatient(c)
}

// DeletePatient godoc
// @Summary Delete a patient
// @Description Delete a patient
// @Tags patients
// @ID delete-patient
// @Accept json
// @Produce json
// @Param id path string true "Patient ID"
// @Router /patient/{id} [delete]
func DeletePatient(c *gin.Context) {
	services.DeletePatient(c)
}
