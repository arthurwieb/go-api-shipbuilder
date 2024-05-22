package handlers

import (
	"goshipbuilder/database"
	"goshipbuilder/models"

	"github.com/gofiber/fiber/v2"
)

func ListShips(c *fiber.Ctx) error {
	ships := []models.Ship{}
	database.DB.Preload("ShipType").Preload("UpgradeOutfitting").Preload("UpgradeSpecialty").Preload("UpgradeSupplies").Find(&ships)

	return c.Status(200).JSON(ships)
}

func ShowShipById(c *fiber.Ctx) error {
	id := c.Params("id")
	ship := models.Ship{}
	if err := database.DB.Preload("ShipType").Preload("UpgradeOutfitting").Preload("UpgradeSpecialty").Preload("UpgradeSupplies").Where("id = ?", id).First(&ship).Error; err != nil {
		return c.Status(404).SendString("Ship not found")
	}
	return c.Status(200).JSON(ship)
}

func CreateShip(c *fiber.Ctx) error {
	ship := new(models.Ship)
	// Parse the request body into the ship struct
	if err := c.BodyParser(ship); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request body",
			"error":   err.Error(),
		})
	}
	// Attempt to create the ship in the database
	result := database.DB.Create(&ship)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating ship",
			"error":   result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ship)
}

func DeleteShipById(c *fiber.Ctx) error {
	id := c.Params("id")
	ship := models.Ship{}
	if err := database.DB.Where("id = ?", id).First(&ship).Error; err != nil {
		return c.Status(404).JSON("Ship not found")
	}
	if err := database.DB.Delete(&ship).Error; err != nil {
		return c.Status(500).JSON("Failed to delete the ship")
	}

	return c.Status(204).JSON("Ship deleted successfully")
}
