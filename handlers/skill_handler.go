package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"wannn-site-rebuild-api/config"
	"wannn-site-rebuild-api/models"
)

type CreateSkillCategoryRequest struct {
	Title  string   `json:"title"`
	Skills []string `json:"skills"`
}

func GetSkillCategories(c *fiber.Ctx) error {
	var categories []models.SkillCategory
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Convert each category's skills back to array
	type SkillCategoryResponse struct {
		ID     uint     `json:"id"`
		Title  string   `json:"title"`
		Skills []string `json:"skills"`
	}

	var response []SkillCategoryResponse
	for _, cat := range categories {
		var skills []string
		json.Unmarshal([]byte(cat.Skills), &skills)
		response = append(response, SkillCategoryResponse{
			ID:     cat.ID,
			Title:  cat.Title,
			Skills: skills,
		})
	}

	return c.JSON(response)
}

func GetSkillCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.SkillCategory
	result := config.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Skill category not found",
		})
	}

	// Convert skills back to array
	var skills []string
	json.Unmarshal([]byte(category.Skills), &skills)

	return c.JSON(fiber.Map{
		"id":     category.ID,
		"title":  category.Title,
		"skills": skills,
	})
}

func CreateSkillCategory(c *fiber.Ctx) error {
	var req CreateSkillCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Convert skills array to JSON string
	skillsJSON, err := json.Marshal(req.Skills)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process skills",
		})
	}

	category := models.SkillCategory{
		Title:  req.Title,
		Skills: string(skillsJSON),
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Convert back to response format
	var skills []string
	json.Unmarshal([]byte(category.Skills), &skills)
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":     category.ID,
		"title":  category.Title,
		"skills": skills,
	})
}

func UpdateSkillCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.SkillCategory
	
	// Check if category exists
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Skill category not found",
		})
	}

	var req CreateSkillCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert skills array to JSON string
	skillsJSON, err := json.Marshal(req.Skills)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process skills",
		})
	}

	category.Title = req.Title
	category.Skills = string(skillsJSON)

	config.DB.Save(&category)

	// Convert back to response format
	var skills []string
	json.Unmarshal([]byte(category.Skills), &skills)

	return c.JSON(fiber.Map{
		"id":     category.ID,
		"title":  category.Title,
		"skills": skills,
	})
}

func DeleteSkillCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.SkillCategory{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Skill category not found",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
} 