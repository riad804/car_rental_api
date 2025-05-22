package config

import (
	"fmt"
	"log"
	"os"

	"github.com/car_rental_api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		// If .env file doesn't exist, try app.env
		err = godotenv.Load("app.env")
		if err != nil {
			log.Printf("Warning: .env file not found. Using environment variables.")
		}
	}

	// Verify required environment variables
	requiredEnvVars := []string{
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_PORT",
		"JWT_SECRET",
	}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Printf("Warning: %s environment variable is not set", envVar)
		}
	}
}

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate models
	err = db.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Rental{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
