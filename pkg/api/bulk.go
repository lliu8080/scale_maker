package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// createResourcesFromBody creates all the resources passed via request body.
//
//	@Summary		Creates all the resources passed via request body.
//	@Description	Creates all the resources passed via request body.
//	@Tags			Bulk API
//	@Accept			application/yaml
//	@Param			body_param	body	string	true	"body_param"
//	@Produce		json
//	@Success		201	"Sample result: "{\"message\":\"k8s resources have been created successfully\",\"status\":201}" string
//	@Router			/api/v1/bulk/create [post]
func createResourcesFromBody(c *fiber.Ctx) error {
	c.Accepts("application/yaml")
	resourceKind := ""
	if err := k8s.CreateResourceFromData(kc, c.Body(), resourceKind); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error: create k8s resources failed with error - " + err.Error() + "!",
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "k8s resources have been created successfully",
	})
}
