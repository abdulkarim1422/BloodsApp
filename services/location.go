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

func GetLocation(c *gin.Context) {
	ip := c.ClientIP()
	// ip := c.Query("ip")
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get location"})
		return
	}
	defer resp.Body.Close()

	var geo GeoResponse
	err = json.NewDecoder(resp.Body).Decode(&geo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"country": geo.Country, "city": geo.City})
}
