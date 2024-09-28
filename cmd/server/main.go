package main

import (
	"log"
	"startup/internal/config"
	"startup/internal/database"
	"startup/internal/router"
	"startup/internal/user"

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

	// Auto-migrate User model
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize and start the router
	r := router.InitRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
