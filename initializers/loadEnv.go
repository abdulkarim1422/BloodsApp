package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PythonCaller string

func python_caller() {
	PythonCaller = os.Getenv("PythonCaller")
	if PythonCaller == "" {
		PythonCaller = "python3"
	}
}

func LoadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}
	python_caller()
}
