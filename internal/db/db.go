package db

import (
	"fmt"
	"os"

	"github.com/ankan-withadream/kiGo/src/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database

	fmt.Println("DB connected successfully")

	if err = DB.AutoMigrate(&models.Chatroom{}); err != nil {
		panic("Failed to migrate database")
	}

	return DB
}

func Get() *gorm.DB {

	return DB
}
