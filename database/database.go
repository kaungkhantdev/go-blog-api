package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"

	iconModel "go-blog-api/internal/icon/models"
	otpModel "go-blog-api/internal/otp/models"
	tagModel "go-blog-api/internal/tag/models"
	userModel "go-blog-api/internal/user/models"
)

var (
	dbname   = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	DB       *gorm.DB
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Log successful connection
	log.Println("Successfully connected to the database.")

	// Run AutoMigrate
	autoMigrateDB := DB.AutoMigrate(&userModel.User{}, &otpModel.Otp{}, &iconModel.Icon{}, &tagModel.Tag{})
	if err := autoMigrateDB; err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	log.Println("Database migration completed successfully.")
	return DB, nil
}
