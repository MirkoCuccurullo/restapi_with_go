package main 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	//home
	app.Get("/", handlers.Home)

	//facts
	app.Get("/listFacts", handlers.ListFacts)
	app.Post("/addFacts", handlers.CreateFact)
	app.Delete("/facts/:id", handlers.DeleteFact)
	app.Put("/facts/:id", handlers.UpdateFact)

	//caps
	app.Get("/listCaps", handlers.ListCaps)
	app.Post("/addCaps", handlers.CreateCaps)
	app.Delete("/caps/:id", handlers.DeleteCaps)
	app.Put("/caps/:id", handlers.UpdateCaps)
}
