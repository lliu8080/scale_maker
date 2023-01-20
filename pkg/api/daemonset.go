package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listDaemonsets doc
func listDaemonsets(c *fiber.Ctx) error {
	resource := "daemonsets"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "apps", "v1", resource, namespace)
}
