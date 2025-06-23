package model

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Name     string
	Position string
	Points   string
}

type Constructor struct {
	gorm.Model
	Name       string
	Nationality string
	Wins       int
}