package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// createResourcesFromBody creates all the resources passed via request body.
// @Summary Creates all the resources passed via request body.
// @Description Creates all the resources passed via request body.
// @Tags Bulk Kubernetes
// @Accept  json
// @Produce  json
// @Success 200 "Sample result: "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}"" string
// @Router /api/v1/bulk/create [post]
func createResourcesFromBody(c *fiber.Ctx) error {
	if err := k8s.CreateReourceFromData(kc, c.Body()); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error: create workload(s) failed with error " + err.Error() + "!",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "workload(s) have been created successfully",
	})
}
