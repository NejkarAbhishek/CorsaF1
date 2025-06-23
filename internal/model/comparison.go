package model

type ComparisonResult struct {
	Season     string `json:"season"`
	DriverA    string `json:"driver_a"`
	DriverB    string `json:"driver_b"`
	WinsA      int    `json:"wins_a"`
	WinsB      int    `json:"wins_b"`
}

// File: internal/repository/db.go
package repository

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"corsaf1/internal/model"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	db.AutoMigrate(&model.Driver{}, &model.Constructor{})
	DB = db
}

func SaveDrivers(drivers []model.Driver) error {
	return DB.Create(&drivers).Error
}

func SaveConstructors(constructors []model.Constructor) error {
	return DB.Create(&constructors).Error
}
