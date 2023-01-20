package api

import "github.com/gofiber/fiber/v2"

// listServices doc
func listServices(c *fiber.Ctx) error {
	resource := "services"
	namespace := c.Query("namespace")
	return listResources(c, "core", "v1", resource, namespace)
}
