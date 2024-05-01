package handlers

import (
	"goshipbuilder/database"
	"goshipbuilder/models"

	"github.com/gofiber/fiber/v2"
)

func ListUpgradesOutfitting(c *fiber.Ctx) error {
	upgradeOutfitting := []models.UpgradeOutfitting{}
	database.DB.Find(&upgradeOutfitting)

	return c.Status(200).JSON(upgradeOutfitting)
}

func GetUpgradeOutfittingById(c *fiber.Ctx) error {
	id := c.Params("id")
	upgradeOutfitting := models.UpgradeOutfitting{}
	if err := database.DB.Where("id = ?", id).First(&upgradeOutfitting).Error; err != nil {
		return c.Status(404).SendString("UpgradeOutfitting not found")
	}

	return c.Status(200).JSON(upgradeOutfitting)
}
