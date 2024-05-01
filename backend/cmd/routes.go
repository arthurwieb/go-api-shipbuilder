package main

import (
	"goshipbuilder/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Use(cors.New())
	//facts
	app.Get("/", handlers.ListFacts)
	app.Get("/fact/:id", handlers.ShowFactById)
	app.Post("/fact", handlers.CreateFact)
	app.Delete("/fact/:id", handlers.DeleteFactById)

	//ship types
	app.Get("/shiptypes", handlers.ListShipTypes)
	app.Get("/shiptypes/:id", handlers.ShowShipTypeById)

	// upgrade outfitting
	app.Get("/upgradeoutfitting", handlers.ListUpgradesOutfitting)
	app.Get("/upgradeoutfitting/:id", handlers.GetUpgradeOutfittingById)
}
