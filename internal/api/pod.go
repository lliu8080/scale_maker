package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/internal/util"
)

// ListPod doc
func ListPod(c *fiber.Ctx) error {
	namespace := c.Query("namespace")
	if namespace == "" {
		namespace = "default"
	}
	if k8sClients.dynamicClient == nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error initialize k8s client!",
			},
		)
	}

	list, err := util.ListDynamicK8SObjectByNames(
		k8sClients.ctx,
		k8sClients.dynamicClient,
		"",
		"v1",
		"pod",
		namespace,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error getting result when trying to list pods!",
			},
		)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":         http.StatusOK,
		"namespace":      namespace,
		"number_of_pods": len(list),
		"pods":           list,
	})
}
