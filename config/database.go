package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetupDatabase() {
	initErr := godotenv.Load()
	if initErr != nil {
		log.Fatal(".env file not found")
	}

	var config gorm.Config
	var err error

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)

	if os.Getenv("ENABLE_GORM_LOGGER") != "" {
		config = gorm.Config{}
	} else {
		config = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	DB, err = gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect database")
	}

	fmt.Println("Database Connected Successfully")
}
