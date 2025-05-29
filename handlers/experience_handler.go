package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"wannn-site-rebuild-api/config"
	"wannn-site-rebuild-api/models"
)

type CreateExperienceRequest struct {
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Period      string   `json:"period"`
	Description []string `json:"description"`
}

func GetExperiences(c *fiber.Ctx) error {
	var experiences []models.Experience
	result := config.DB.Find(&experiences)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Convert each experience's description back to array
	type ExperienceResponse struct {
		ID          uint     `json:"id"`
		Title       string   `json:"title"`
		Company     string   `json:"company"`
		Period      string   `json:"period"`
		Description []string `json:"description"`
	}

	var response []ExperienceResponse
	for _, exp := range experiences {
		var desc []string
		json.Unmarshal([]byte(exp.Description), &desc)
		response = append(response, ExperienceResponse{
			ID:          exp.ID,
			Title:       exp.Title,
			Company:     exp.Company,
			Period:      exp.Period,
			Description: desc,
		})
	}

	return c.JSON(response)
}

func GetExperienceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var experience models.Experience
	result := config.DB.First(&experience, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Experience not found",
		})
	}

	// Convert description back to array
	var desc []string
	json.Unmarshal([]byte(experience.Description), &desc)

	return c.JSON(fiber.Map{
		"id":          experience.ID,
		"title":       experience.Title,
		"company":     experience.Company,
		"period":      experience.Period,
		"description": desc,
	})
}

func CreateExperience(c *fiber.Ctx) error {
	var req CreateExperienceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Convert description array to JSON string
	descJSON, err := json.Marshal(req.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process description",
		})
	}

	experience := models.Experience{
		Title:       req.Title,
		Company:     req.Company,
		Period:      req.Period,
		Description: string(descJSON),
	}

	result := config.DB.Create(&experience)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Convert back to response format
	var desc []string
	json.Unmarshal([]byte(experience.Description), &desc)
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":          experience.ID,
		"title":       experience.Title,
		"company":     experience.Company,
		"period":      experience.Period,
		"description": desc,
	})
}

func UpdateExperience(c *fiber.Ctx) error {
	id := c.Params("id")
	var experience models.Experience
	
	// Check if experience exists
	if err := config.DB.First(&experience, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Experience not found",
		})
	}

	var req CreateExperienceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert description array to JSON string
	descJSON, err := json.Marshal(req.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process description",
		})
	}

	experience.Title = req.Title
	experience.Company = req.Company
	experience.Period = req.Period
	experience.Description = string(descJSON)

	config.DB.Save(&experience)

	// Convert back to response format
	var desc []string
	json.Unmarshal([]byte(experience.Description), &desc)

	return c.JSON(fiber.Map{
		"id":          experience.ID,
		"title":       experience.Title,
		"company":     experience.Company,
		"period":      experience.Period,
		"description": desc,
	})
}

func DeleteExperience(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.Experience{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Experience not found",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
} 