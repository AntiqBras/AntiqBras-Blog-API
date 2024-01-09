package domain

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"isaacszf.antiqbrasblog.com/domain/models"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	err = database.AutoMigrate(&models.Writer{}, &models.Post{})
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	// Default Writer
	var writer models.Writer
	result := database.Where("username = ?", "writer").First(&writer)

	// If not found, create the default writer
	if result.Error == gorm.ErrRecordNotFound {
		defaultWriter := models.Writer{
			Username: os.Getenv("DEFAULT_WRITER_NAME"),
			Password: os.Getenv("DEFAULT_WRITER_PASSWORD"),
			Author: os.Getenv("DEFAULT_WRITER_AUTHOR"),
		}

		err = database.Create(&defaultWriter).Error
		if err != nil {
			log.Fatal("Failed to create default writer: ", err)
		}
	}

	DB = database
}
