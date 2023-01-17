package api

import (
	"github.com/gofiber/fiber/v2"
)

// listPod doc
func listPod(c *fiber.Ctx) error {
	resource := "pods"
	namespace := c.Query("namespace")
	return listResources(c, "", "v1", resource, namespace)
}
