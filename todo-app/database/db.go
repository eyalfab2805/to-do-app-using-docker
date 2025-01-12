package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// getEnvOrPanic retrieves the value of an environment variable or exits if not set.
func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

// InitDB initializes the database connection with retries.
func InitDB() {
	dbHost := getEnvOrPanic("DB_HOST")
	dbPort := getEnvOrPanic("DB_PORT")
	dbUser := getEnvOrPanic("DB_USER")
	dbPassword := getEnvOrPanic("DB_PASSWORD")
	dbName := getEnvOrPanic("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	for retries := 0; retries < 5; retries++ {
		DB, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			log.Println("Connected to the database successfully.")
			return
		}
		log.Printf("Failed to connect to the database. Retrying in 2 seconds... (%d/5)", retries+1)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Failed to connect to the database after multiple attempts: %v", err)
}

// CloseDB gracefully closes the database connection.
func CloseDB() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}
