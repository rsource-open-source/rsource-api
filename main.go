package main

import (
	"rsource-api/rtime"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/time:id", rtime.GetTime)
}

func main() {

	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(3000)
}
