package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
)

var unverifiedDonors = make(map[string]models.Donor)

func CreateDonor(c *gin.Context) {
	var newDonor models.Donor
	if err := c.ShouldBind(&newDonor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle the CarAvailable checkbox value
	newDonor.CarAvailable = c.PostForm("CarAvailable") == "true"

	// Generate and assign verification code
	newDonor.Verify = generateVerificationCode()

	// Send verification code via WhatsApp
	go sendVerificationCode(newDonor.PhoneNumber, newDonor.Verify)

	// Temporarily store the donor data
	tempID := strconv.Itoa(len(unverifiedDonors) + 1)
	unverifiedDonors[tempID] = newDonor

	c.JSON(http.StatusCreated, gin.H{"id": tempID})
}

func VerifyDonor(c *gin.Context) {
	donorID := c.Param("id")
	var request struct {
		VerificationCode string `json:"verificationCode"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	donor, exists := unverifiedDonors[donorID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}

	if donor.Verify != request.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code"})
		return
	}

	// Verification successful, clear the verification code and add to donors
	donor.Verify = ""

	// Create a new donor in the database
	create_err := repositories.CreateDonor(&donor)
	if create_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": create_err.Error()})
		return
	}

	// Remove from unverified donors after successful verification and commit
	delete(unverifiedDonors, donorID)

	c.JSON(http.StatusOK, gin.H{"message": "Verification successful"})
}

func DonorByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	d, err := repositories.GetDonorByID(intID)
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

func UpdateDonor(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var updatedDonor models.Donor
	if err := c.ShouldBind(&updatedDonor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedDonor.ID = uint(intID)
	err = repositories.UpdateDonor(&updatedDonor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Donor updated successfully")
	c.Redirect(http.StatusFound, "/donors")
}

func DeleteDonor(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	err = repositories.DeleteDonor(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Donor deleted successfully"})
}
