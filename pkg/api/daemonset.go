package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listDaemonsets gets the list of the daemonsets in the k8s cluster.
// @Summary Gets the list of the Daemonsets in the k8s cluster.
// @Description Gets the list of the daemonsets in the k8s cluster.
// @Tags Kubernetes
// @Accept  json
// @Produce  json
// @Success 200 "Sample result: "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}"" string
// @Router /api/v1/daemonset/list [get]
func listDaemonsets(c *fiber.Ctx) error {
	resource := "daemonsets"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "apps", "v1", resource, namespace)
}
