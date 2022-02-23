package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/rsource-open-source/rsource-api/database"

	// "github.com/rsource-open-source/rsource-api/discord"
	v0 "github.com/rsource-open-source/rsource-api/routes/v0"
)

func setupRoutes(app *fiber.App) {
	appv0 := app.Group("/v0")

	appv0.Post("/redirect", v0.AddRedirect)
	// appv0.Get("redirect/:id", v0.GetRedirect)
	// app.Get("/v0/redirect", routes.GetRedirects)
	appv0.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("v0 base")
	})
}

func main() {

	// discord.StartBot()

	// database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("404")
	})

	log.Fatal(app.Listen(":3000"))
}
