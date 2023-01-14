package api

import "github.com/gofiber/fiber/v2"

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	// return c.Status(404).SendFile("./assets/static/private/404.html")
	return c.Status(404).JSON(fiber.Map{
		"status": "404 page not found!",
	})
}
