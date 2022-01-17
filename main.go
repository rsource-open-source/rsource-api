package main

import (
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

// func setupRoutes(app *fiber.App) {
// 	app.Get("/v1/time:id", rtime.GetTime)
// }

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(3000)
}
