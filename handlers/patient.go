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
// @Router /patients/{id} [get]
func PatientByID(c *gin.Context) {
	services.PatientByID(c)
}

// CreatePatient godoc
// @Summary Create a patient
// @Description Create a patient
// @Tags patients
// @ID create-patient
// @Accept json
// @Produce json
// @Param patient body models.Patient true "Patient object"
// @Router /patients [post]
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
// @Router /patients/{id} [put]
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
// @Router /patients/{id} [delete]
func DeletePatient(c *gin.Context) {
	services.DeletePatient(c)
}
