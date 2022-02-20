package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rsource-open-source/rsource-api/database"
	v0 "github.com/rsource-open-source/rsource-api/routes/v0"
)

func setupRoutes(app *fiber.App) {
	appv0 := app.Group("/v0")

	appv0.Get("/redirect\\?user_id=:user_id&place_id=:place_id&server_id=:server_id&request_id=:request_id", v0.AddRedirect)
	// app.Get("/v0/redirect", routes.GetRedirects)
	appv0.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("v0 base")
	})
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})

	log.Fatal(app.Listen(":3000"))
}
