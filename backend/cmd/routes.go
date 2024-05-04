package main

import (
	"goshipbuilder/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Use(cors.New())
	//ships
	app.Get("/ships", handlers.ListShips)
	app.Get("/ships/:id", handlers.ShowShipById)
	app.Post("/ships", handlers.CreateShip)
	app.Delete("/ships/:id", handlers.DeleteShipById)

	//ship types
	app.Get("/shiptypes", handlers.ListShipTypes)
	app.Get("/shiptypes/:id", handlers.ShowShipTypeById)

	// upgrade outfitting
	app.Get("/upgrade/outfitting", handlers.ListUpgradesOutfitting)
	app.Get("/upgrade/outfitting/:id", handlers.GetUpgradeOutfittingById)

	// upgrade specialty
	app.Get("/upgrade/specialty", handlers.ListUpgradesSpecialty)
	app.Get("/upgrade/specialty/:id", handlers.GetUpgradeSpecialtyById)

	// upgrade supplies
	app.Get("/upgrade/supplies", handlers.ListUpgradesSupplies)
	app.Get("/upgrade/supplies/:id", handlers.GetUpgradeSuppliesById)
}
