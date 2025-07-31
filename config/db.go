package config

import (
	"fmt"
	"internship-stikom/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env failed")
	}
	fmt.Println("ENV loaded")
	dsn := fmt.Sprintf(
		"host=%s user=%s password =%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed : ", err)
	}
	DB = db
	err = db.AutoMigrate((&models.User{}), &models.Profile{}, &models.Project{}, &models.Internship{})
	if err != nil {
		log.Fatal("Migration Failed : ", err)
	}

	fmt.Println("Connected to database successfully")
}
