package main

import (
	"log"
	"startup/internal/config"
	"startup/internal/database"
	"startup/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load additional configuration
	config.LoadConfig()

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize and start the router
	r := router.InitRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
