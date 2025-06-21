package repository

import (
	"fmt"
	"log"
	"os"

	"CORSAF1/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	err = db.AutoMigrate(&model.Driver{})
	if err != nil {
		log.Fatal("Failed to migrate DB schema: ", err)
	}

	DB = db
}

func SaveDrivers(drivers []model.Driver) error {
	return DB.Create(&drivers).Error
}
