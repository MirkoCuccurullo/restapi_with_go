package main 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/facts", handlers.CreateFact)
}
