package api

import (
	"github.com/gofiber/fiber/v2"
)

// getStatus gets the current status of the application.
//	@Summary		Fetch the current status of the application.
//	@Description	Fetch the current status of the application.
//	@Tags			Status
//	@Accept			json
//	@Produce		json
//	@Success		200	"Sample result: {\"status\":\"alive\"}"	string
//	@Router			/api/v1/ping [get]
func getStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "alive",
	})
}
