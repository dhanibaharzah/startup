package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	// Build the connection string
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully!")
	DB = db
	return DB, nil
}
