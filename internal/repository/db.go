package repository

import (
	"CorsaF1/internal/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	log.Printf("Attempting DB connection with: User=%s, Pass=%s (hidden), DBName=%s\n", dbUser, dbPass, dbName)

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		dbUser, dbPass, dbName)

	log.Printf("Full DSN string: %s\n", dsn) // Print the generated DSN

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Waiting for DB... (%d/10) Error: %v\n", i+1, err) // Log the error during retry
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to DB after retries: %v", err)
	}

	log.Println("Successfully connected to the database. Running migrations.")
	DB.AutoMigrate(&model.Driver{}, &model.Constructor{})
	log.Println("Migrations completed.")
}

func SaveDrivers(drivers []model.Driver) error {
	return DB.Create(&drivers).Error
}

func SaveConstructors(constructors []model.Constructor) error {
	return DB.Create(&constructors).Error
}
