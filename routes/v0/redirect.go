package v0

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rsource-open-source/rsource-api/blueprints"
	"github.com/rsource-open-source/rsource-api/database"
)

func AddRedirect(c *fiber.Ctx) error {

	user_id := c.Params("user_id")
	place_id := c.Params("place_id")
	server_id := c.Params("server_id")
	request_id := c.Params("request_id")

	fmt.Println(user_id, place_id, server_id, request_id)

	redirect := new(blueprints.Redirect)

	// user_id
	uid_int, err := strconv.Atoi(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Unable to convert user_id to int")
	}

	// place_id
	pid_int, err := strconv.Atoi(place_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Unable to convert place_id to int")
	}

	redirect = &blueprints.Redirect{
		UserID:    uid_int,
		PlaceID:   pid_int,
		ServerID:  server_id,
		RequestID: request_id,
	}

	if err := c.BodyParser(redirect); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// database.DB.Db.Create(&redirect)

	return c.Status(fiber.StatusCreated).JSON(redirect)
}

func GetRedirects(c *fiber.Ctx) error {
	redirects := []blueprints.Redirect{}
	database.DB.Db.Find(&redirects)

	if redirects != nil {
		fmt.Println(reflect.TypeOf(redirects))
		return c.Status(fiber.StatusCreated).JSON(redirects)
	} else {
		return c.Status(fiber.StatusNoContent).SendString("No redirects found")
	}
}
