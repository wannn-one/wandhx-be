package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"wannn-site-rebuild-api/config"
	"wannn-site-rebuild-api/models"
)

type CreateProjectRequest struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	Link         string   `json:"link"`
}

func GetProjects(c *fiber.Ctx) error {
	var projects []models.Project
	result := config.DB.Find(&projects)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Convert each project's technologies back to array
	type ProjectResponse struct {
		ID           uint     `json:"id"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Technologies []string `json:"technologies"`
		Link         string   `json:"link"`
	}

	var response []ProjectResponse
	for _, proj := range projects {
		var tech []string
		json.Unmarshal([]byte(proj.Technologies), &tech)
		response = append(response, ProjectResponse{
			ID:           proj.ID,
			Title:        proj.Title,
			Description:  proj.Description,
			Technologies: tech,
			Link:         proj.Link,
		})
	}

	return c.JSON(response)
}

func GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var project models.Project
	result := config.DB.First(&project, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	// Convert technologies back to array
	var tech []string
	json.Unmarshal([]byte(project.Technologies), &tech)

	return c.JSON(fiber.Map{
		"id":           project.ID,
		"title":        project.Title,
		"description":  project.Description,
		"technologies": tech,
		"link":         project.Link,
	})
}

func CreateProject(c *fiber.Ctx) error {
	var req CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Convert technologies array to JSON string
	techJSON, err := json.Marshal(req.Technologies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process technologies",
		})
	}

	project := models.Project{
		Title:        req.Title,
		Description:  req.Description,
		Technologies: string(techJSON),
		Link:         req.Link,
	}

	result := config.DB.Create(&project)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Convert back to response format
	var tech []string
	json.Unmarshal([]byte(project.Technologies), &tech)
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":           project.ID,
		"title":        project.Title,
		"description":  project.Description,
		"technologies": tech,
		"link":         project.Link,
	})
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var project models.Project
	
	// Check if project exists
	if err := config.DB.First(&project, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}

	var req CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert technologies array to JSON string
	techJSON, err := json.Marshal(req.Technologies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process technologies",
		})
	}

	project.Title = req.Title
	project.Description = req.Description
	project.Technologies = string(techJSON)
	project.Link = req.Link

	config.DB.Save(&project)

	// Convert back to response format
	var tech []string
	json.Unmarshal([]byte(project.Technologies), &tech)

	return c.JSON(fiber.Map{
		"id":           project.ID,
		"title":        project.Title,
		"description":  project.Description,
		"technologies": tech,
		"link":         project.Link,
	})
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.Project{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Project not found",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
} 