package handlers

import (
	"goshipbuilder/database"
	"goshipbuilder/models"

	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func ShowFactById(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	if err := database.DB.Db.Where("id = ?", id).First(&fact).Error; err != nil {
		return c.Status(404).SendString("Fact not found")
	}
	return c.Status(200).JSON(fact)
}

func DeleteFactById(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	if err := database.DB.Db.Where("id = ?", id).First(&fact).Error; err != nil {
		return c.Status(404).JSON("Fact not found")
	}
	// Delete the fact
	if err := database.DB.Db.Delete(&fact).Error; err != nil {
		return c.Status(500).JSON("Failed to delete the fact")
	}

	return c.Status(204).JSON("Fact deleted successfully")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}
