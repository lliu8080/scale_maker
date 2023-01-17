package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	// return c.Status(404).SendFile("./assets/static/private/404.html")
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"status": "Error: 404 page not found!",
	})
}
