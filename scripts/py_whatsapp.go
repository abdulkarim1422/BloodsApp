package scripts

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func PyWhatsapp(c *gin.Context) {
	// Define the Python script and arguments
	pythonScript := "whatsapp.py"
	dataFile := "body.json"
	messageFile := "message.txt"
	xCord := "958"
	yCord := "968"

	// Create the command
	cmd := exec.Command("python", pythonScript, dataFile, messageFile, xCord, yCord)

	// Run the command and capture the output and error
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// Print the output
	fmt.Printf("Output: %s\n", output)
}
