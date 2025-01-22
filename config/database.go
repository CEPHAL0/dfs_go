package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var once sync.Once

func Initialize(db *gorm.DB) {
	once.Do(func() {
		DB = db
	})
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("Database not initialized. Call database.Initialize(db) first.")
	}
	return DB
}

// WithTransaction manages a database transaction
func WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

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
