package rtime

import (
	"github.com/gofiber/fiber"
)

func GetTime(c *fiber.Ctx) {
	c.Send("time")
}
