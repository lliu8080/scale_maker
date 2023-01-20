package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listDeployments doc
func listDeployments(c *fiber.Ctx) error {
	resource := "deployments"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "apps", "v1", resource, namespace)
}
