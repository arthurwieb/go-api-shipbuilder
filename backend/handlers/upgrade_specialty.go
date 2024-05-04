package handlers

import (
	"goshipbuilder/database"
	"goshipbuilder/models"

	"github.com/gofiber/fiber/v2"
)

func ListUpgradesSpecialty(c *fiber.Ctx) error {
	upgradeSpecialty := []models.UpgradeSpecialty{}
	database.DB.Find(&upgradeSpecialty)

	return c.Status(200).JSON(upgradeSpecialty)
}

func GetUpgradeSpecialtyById(c *fiber.Ctx) error {
	id := c.Params("id")
	upgradeSpecialty := models.UpgradeSpecialty{}
	if err := database.DB.Where("id = ?", id).First(&upgradeSpecialty).Error; err != nil {
		return c.Status(404).SendString("UpgradeSpecialty not found")
	}

	return c.Status(200).JSON(upgradeSpecialty)
}
