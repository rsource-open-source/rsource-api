package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rsource-open-source/rsource-api/database"
	"github.com/rsource-open-source/rsource-api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/v1/professor", routes.AddProfessor)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON("Not found")
	})

	log.Fatal(app.Listen(":3000"))
}
