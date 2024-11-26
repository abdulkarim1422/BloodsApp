package handlers

import (
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
// @Router /schedualed-requests [get]
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
