package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func DbConn() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	// Get database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}

func GetQuery(query string) (map[string][]string, error) {

	compatibleBloodTypes := make(map[string][]string)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bloodType string
		var compatibleTypes string

		err := rows.Scan(&bloodType, &compatibleTypes)
		if err != nil {
			return nil, err
		}

		typesArray := strings.Split(compatibleTypes, ",")
		compatibleBloodTypes[bloodType] = typesArray
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return compatibleBloodTypes, nil
}
