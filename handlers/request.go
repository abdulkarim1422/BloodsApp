package handlers

import (
	"net/http"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/services"
	"github.com/gin-gonic/gin"
)

// CreateSchedualedRequest godoc
// @Summary Create a schedualed request
// @Description Create a schedualed request
// @Tags request
// @Accept  json
// @Produce  json
// @Param request body models.SchedualedRequest true "SchedualedRequest"
// @Success 200 {string} string "Schedualed request created"
// @Router /scheduale-request [post]
func CreateSchedualedRequest(c *gin.Context) {
	services.CreateSchedualedRequest(c)
}

// GetAllSchedualedRequests godoc
// @Summary Get all schedualed requests
// @Description Get all schedualed requests
// @Tags request
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SchedualedRequest "List of schedualed requests"
// @Router /get-all-schedualed-requests [get]
func GetAllSchedualedRequests(c *gin.Context) {
	services.GetAllSchedualedRequests(c)
}

// DeleteScheduledRequest godoc
// @Summary Delete a schedualed request
// @Description Delete a schedualed request
// @Tags request
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the request"
// @Success 200 {string} string "Request deleted"
// @Router /schedualed-request/{id} [delete]
func DeleteScheduledRequest(c *gin.Context) {
	services.DeleteScheduledRequest(c)
}

// PerformSchedualedRequest godoc
// @Summary Perform a schedualed request
// @Description Perform a schedualed request
// @Tags request
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the request"
// @Success 200 {string} string "Schedualed request performed successfully"
// @Router /perform-schedualed-request/{id} [post]
func PerformSchedualedRequest(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := services.PerformSchedualedRequest(intID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedualed request performed successfully"})
}

func MarkAsDonated(c *gin.Context) {
	var donationForm struct {
		RequestID string `form:"requestID" binding:"required"`
		Type      string `form:"type" binding:"required"`
		Feedback  string `form:"feedback"`
	}

	if err := c.ShouldBind(&donationForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestIDint, err := strconv.Atoi(donationForm.RequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := services.MarkAsDonated(requestIDint, donationForm.Type, donationForm.Feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request marked as donated and donor feedback saved, feedback will be sent to the patient"})
}

func MarkAsReceived(c *gin.Context) {
	var donationForm struct {
		RequestID string `form:"requestID" binding:"required"`
		QQQ       string `form:"qqq" binding:"required"`
		Feedback  string `form:"feedback"`
	}

	requestIDint, err := strconv.Atoi(donationForm.RequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := services.MarkAsReceived(requestIDint, donationForm.QQQ, donationForm.Feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request marked as received and patient feedback saved, feedback will be sent to the donor"})
}
