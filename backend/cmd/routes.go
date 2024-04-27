package main

import (
	"goshipbuilder/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Get("/fact/:id", handlers.ShowFactById)
	app.Post("/fact", handlers.CreateFact)
	app.Delete("/fact/:id", handlers.DeleteFactById)
}
