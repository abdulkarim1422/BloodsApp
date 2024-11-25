package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/abdulkarim1422/BloodsApp/initializers"
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

	// Change the working directory to scripts
	scriptsDir, err := filepath.Abs("scripts")
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get absolute path"})
		return
	}

	// Run the PyWhatsapp function
	cmd := exec.Command(initializers.PythonCaller, "whatsapp_match.py", "body.json", "message.txt")
	cmd.Dir = scriptsDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running PyWhatsapp: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to run PyWhatsapp: %s", err)})
		return
	}

	// Log the output of the command
	fmt.Printf("Command output: %s\n", output)

	c.JSON(http.StatusOK, gin.H{"message": "تمّ انتهاء عمل البرنامج بنجاح", "output": string(output)})
}

func generateVerificationCode() string {
	rand.Seed(int64(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func sendVerificationCode(phone, code string) {
	cmd := exec.Command(initializers.PythonCaller, "scripts/whatsapp_verify.py", phone, code)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error sending verification code: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		return
	}
	fmt.Printf("Verification code sent successfully: %s\n", output)
}

func SendWhatsappMessage(phone string, message string, txt_file_name string) {
	cmd := exec.Command(initializers.PythonCaller, "whatsapp_message.py", phone, message, txt_file_name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error sending WhatsApp message: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		return
	}
	fmt.Printf("WhatsApp message sent successfully: %s\n", output)

}
