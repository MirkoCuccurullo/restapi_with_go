package main 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/database"
)

func main() {

	database.ConnectDb()

	app := fiber.New()
	
	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Trivia App")
	})

	app.Listen(":3000")
}