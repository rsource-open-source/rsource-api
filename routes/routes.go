package routes

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/rsource-open-source/rsource-api/blueprints"
	"github.com/zeimedee/go-postgres/database"
)

// func AddProfessor(c *fiber.Ctx) error {
// 	professor := new(models.Professor)
// 	if err := c.BodyParser(professor); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	database.DB.Db.Create(&professor)

// 	return c.Status(201).JSON(professor)
// }

func AddRedirect(c *fiber.Ctx) error {
	redirect := new(models.Redirect)
	if err := c.BodyParser(redirect); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&redirect)

	return c.Status(201).JSON(redirect)
}
