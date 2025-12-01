package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load the .env file")
	}

	dns := os.Getenv("DB_URL")

	db, errs := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if errs != nil {
		log.Fatal("unable to connect the DB")
	}

	return db
}
