package db

import (
	"goproj2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase() {
	connStr := "host=db user=postgres password=postgres dbname=postgres port=5432"
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})

	DB = database
}