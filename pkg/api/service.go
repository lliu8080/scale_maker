package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listServices doc
func listServices(c *fiber.Ctx) error {
	resource := "services"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "core", "v1", resource, namespace)
}
