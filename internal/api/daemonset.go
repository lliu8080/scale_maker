package api

import "github.com/gofiber/fiber/v2"

// listDaemonset doc
func listDaemonsets(c *fiber.Ctx) error {
	resource := "daemonsets"
	namespace := c.Query("namespace")
	return listResources(c, "apps", "v1", resource, namespace)
}
