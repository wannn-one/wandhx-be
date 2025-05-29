package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"wannn-site-rebuild-api/models"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, will use system environment variables")
	}

	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("Database configuration environment variables are required (DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop existing tables
	err = db.Migrator().DropTable(&models.Experience{}, &models.Project{}, &models.SkillCategory{})
	if err != nil {
		log.Fatal("Failed to drop tables:", err)
	}

	// Auto Migrate the schemas
	err = db.AutoMigrate(
		&models.Experience{},
		&models.Project{},
		&models.SkillCategory{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connection established and migrations completed")

	// Seed the database with initial data
	SeedDatabase()
} 