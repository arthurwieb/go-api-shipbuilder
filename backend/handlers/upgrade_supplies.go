package handlers

import (
	"goshipbuilder/database"
	"goshipbuilder/models"

	"github.com/gofiber/fiber/v2"
)

func ListUpgradesSupplies(c *fiber.Ctx) error {
	upgradeSupplies := []models.UpgradeSupplies{}
	database.DB.Find(&upgradeSupplies)

	return c.Status(200).JSON(upgradeSupplies)
}

func GetUpgradeSuppliesById(c *fiber.Ctx) error {
	id := c.Params("id")
	upgradeSupplies := models.UpgradeSupplies{}
	if err := database.DB.Where("id = ?", id).First(&upgradeSupplies).Error; err != nil {
		return c.Status(404).SendString("UpgradeSupplies not found")
	}

	return c.Status(200).JSON(upgradeSupplies)
}
