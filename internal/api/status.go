package api

import (
	"github.com/gofiber/fiber/v2"
)

// UserGet returns a user
func GetStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "alive",
	})
}
