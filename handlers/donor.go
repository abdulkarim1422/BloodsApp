package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// DonorByID is a handler function to get a donor by ID
// @Summary Get a donor by ID
// @Description Get a donor by ID
// @Tags donors
// @ID get-donor-by-id
// @Produce json
// @Param id path int true "Donor ID"
// @Router /donor/{id} [get]
func DonorByID(c *gin.Context) {
	services.DonorByID(c)
}

// CreateDonor is a handler function to create a donor
// @Summary Create a donor
// @Description Create a donor
// @Tags donors
// @ID create-donor
// @Accept json
// @Produce json
// @Param donor body models.Donor true "Donor input"
// @Router /donor [post]
func CreateDonor(c *gin.Context) {
	services.CreateDonor(c)
}

// UpdateDonor is a handler function to update a donor
// @Summary Update a donor
// @Description Update a donor
// @Tags donors
// @ID update-donor
// @Accept json
// @Produce json
// @Param id path int true "Donor ID"
// @Param donor body models.Donor true "Donor input"
// @Router /donor/{id} [put]
func UpdateDonor(c *gin.Context) {
	services.UpdateDonor(c)
}

// DeleteDonor is a handler function to delete a donor
// @Summary Delete a donor
// @Description Delete a donor
// @Tags donors
// @ID delete-donor
// @Param id path int true "Donor ID"
// @Router /donor/{id} [delete]
func DeleteDonor(c *gin.Context) {
	services.DeleteDonor(c)
}
