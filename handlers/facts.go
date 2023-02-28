package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	database.DB.Db.Delete(&fact, id)
	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := models.Fact{}
	database.DB.Db.Find(&fact, id)

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	database.DB.Db.Save(&fact)
	return c.Status(200).JSON(fact)
}


func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}
