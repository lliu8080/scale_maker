package api

import (
	"github.com/gofiber/fiber/v2"
)

// getStatus doc
func getStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "alive",
	})
}
