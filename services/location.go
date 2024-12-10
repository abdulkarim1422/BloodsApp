package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeoResponse struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

func GetLocation(c *gin.Context) (string, string) {
	ip := c.ClientIP()
	// ip := c.Query("ip")
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", ""
	}
	defer resp.Body.Close()

	var geo GeoResponse
	err = json.NewDecoder(resp.Body).Decode(&geo)
	if err != nil {
		return "", ""
	}

	return geo.Country, geo.City
}
