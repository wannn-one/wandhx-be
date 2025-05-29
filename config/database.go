package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"wannn-site-rebuild-api/models"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	if supabaseUrl == "" {
		log.Fatal("SUPABASE_URL is required")
	}

	// Extract database URL from Supabase URL
	// Supabase URL format: https://[project_id].supabase.co
	projectID := strings.Split(strings.Split(supabaseUrl, "//")[1], ".")[0]
	dbHost := fmt.Sprintf("db.%s.supabase.co", projectID)
	dbUser := "postgres"
	dbPassword := os.Getenv("SUPABASE_DB_PASSWORD") // This should be the database password, not the anon key
	dbName := "postgres"
	dbPort := "5432"

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