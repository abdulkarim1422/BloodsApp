package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func SendMatchResult(c *gin.Context) {
	var jsonResponse map[string]interface{}
	if err := c.ShouldBindJSON(&jsonResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Log the received JSON
	fmt.Printf("Received JSON: %v\n", jsonResponse)

	// Save the JSON response to /scripts/body.json
	jsonData, err := json.MarshalIndent(jsonResponse, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	if err := ioutil.WriteFile("scripts/body.json", jsonData, 0644); err != nil {
		fmt.Printf("Error writing JSON to file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write JSON to file"})
		return
	}

	// Run the PyWhatsapp function
	cmd0 := exec.Command("powershell", "pwd")
	output, err := cmd0.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running PyWhatsapp: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to run PyWhatsapp: %s", err)})
		return
	}

	cmd := exec.Command("python", "scripts/whatsapp.py", "scripts/body.json", "scripts/message.txt", "958", "968")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running PyWhatsapp: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to run PyWhatsapp: %s", err)})
		return
	}

	// Log the output of the command
	fmt.Printf("Command output: %s\n", output)

	c.JSON(http.StatusOK, gin.H{"message": "Match result sent successfully", "output": string(output)})
}
