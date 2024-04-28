package handlers

import (
	"goshipbuilder/database"
	"goshipbuilder/models"

	"github.com/gofiber/fiber/v2"
)

func ListShipTypes(c *fiber.Ctx) error {
	shiptypes := []models.ShipType{}
	database.DB.Db.Find(&shiptypes)

	return c.Status(200).JSON(shiptypes)
}

func ShowShipTypeById(c *fiber.Ctx) error {
	id := c.Params("id")
	shiptype := models.ShipType{}
	if err := database.DB.Db.Where("id = ?", id).First(&shiptype).Error; err != nil {
		return c.Status(404).SendString("ShipType not found")
	}
	return c.Status(200).JSON(shiptype)
}
