package handlers

import (
	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/services"
)

func CreateSchedualedRequest(request *models.SchedualedRequest) {
	services.CreateSchedualedRequest(request)
}

func GetAllSchedualedRequests() {
	services.GetAllSchedualedRequests()
}

func DeleteScheduledRequest(id int) {
	services.DeleteScheduledRequest(id)
}
