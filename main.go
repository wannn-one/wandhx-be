package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"wannn-site-rebuild-api/config"
	"wannn-site-rebuild-api/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection and run migrations
	config.InitDatabase()

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Routes
	api := app.Group("/")
	
	// Experience routes
	experiences := api.Group("experiences")
	experiences.Get("/", handlers.GetExperiences)
	experiences.Get("/:id", handlers.GetExperienceByID)
	experiences.Post("/", handlers.CreateExperience)
	experiences.Put("/:id", handlers.UpdateExperience)
	experiences.Delete("/:id", handlers.DeleteExperience)

	// Project routes
	projects := api.Group("projects")
	projects.Get("/", handlers.GetProjects)
	projects.Get("/:id", handlers.GetProjectByID)
	projects.Post("/", handlers.CreateProject)
	projects.Put("/:id", handlers.UpdateProject)
	projects.Delete("/:id", handlers.DeleteProject)

	// Skill Category routes
	skills := api.Group("skills")
	skills.Get("/", handlers.GetSkillCategories)
	skills.Get("/:id", handlers.GetSkillCategoryByID)
	skills.Post("/", handlers.CreateSkillCategory)
	skills.Put("/:id", handlers.UpdateSkillCategory)
	skills.Delete("/:id", handlers.DeleteSkillCategory)

	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Fatal(app.Listen(":" + port))
} 