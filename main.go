package main

import (
	"fmt"
	"log"
	"os"

	"go-study-backend/db"
)

func main() {
	
	if err := loadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	server := NewServer()
	log.Println("Server starting on :8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() error {
	
	requiredEnvVars := []string{
		"MYSQL_DATABASE",
		"MYSQL_USER",
		"MYSQL_PASSWORD",
		"MYSQL_ROOT_PASSWORD",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("required environment variable %s is not set", envVar)
		}
	}

	return nil
}
