package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shahirov/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "James",
		LastName:  "Born",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
