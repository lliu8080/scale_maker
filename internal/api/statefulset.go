package api

import "github.com/gofiber/fiber/v2"

// listStatefulsets doc
func listStatefulsets(c *fiber.Ctx) error {
	resource := "statefulsets"
	namespace := c.Query("namespace")
	return listResources(c, "apps", "v1", resource, namespace)
}
