package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
)

func ListCaps(c *fiber.Ctx) error {
	caps := []models.Caps{}
	database.DB.Db.Find(&caps)
	return c.Status(200).JSON(caps)
}

func CreateCaps(c *fiber.Ctx) error {
	caps := new(models.Caps)

	if err := c.BodyParser(caps); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	database.DB.Db.Create(&caps)
	return c.Status(200).JSON(caps)
}

func DeleteCaps(c *fiber.Ctx) error {
	id := c.Params("id")
	caps := models.Caps{}
	database.DB.Db.Delete(&caps, id)
	return c.Status(200).JSON(caps)
}

func UpdateCaps(c *fiber.Ctx) error {
	id := c.Params("id")
	caps := models.Caps{}
	database.DB.Db.Find(&caps, id)

	if err := c.BodyParser(&caps); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"massage": err.Error(),
		})
	}

	database.DB.Db.Save(&caps)
	return c.Status(200).JSON(caps)
}