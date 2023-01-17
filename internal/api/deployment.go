package api

import (
	"github.com/gofiber/fiber/v2"
)

// listDeployment doc
func listDeployments(c *fiber.Ctx) error {
	resource := "deployments"
	namespace := c.Query("namespace")
	return listResources(c, "apps", "v1", resource, namespace)
}
