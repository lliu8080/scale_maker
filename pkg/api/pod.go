package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listPods gets the list of the pods in the k8s cluster.
//
//	@Summary		Gets the list of the pods in the k8s cluster.
//	@Description	Gets the list of the pods in the k8s cluster.
//	@Tags			Pod
//	@Accept			json
//	@Param			namespace	query	string	false	"pod search by namespace"															Format(string)
//	@Param			label		query	string	false	"search pod by label"																Format(string)
//	@Param			by_item		query	string	false	"set by_item=true to return pod results by item with more details, default false."	Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"namespace\":\"default\",\"number_of_pods\":0,\"pods\":[],\"status\":200}" string
//	@Router			/api/v1/pod/list [get]
func listPods(c *fiber.Ctx) error {
	resource := "pods"
	return k8s.ListResources(c, kc, "", "v1", resource)
}

// createPodFromTemplate creates the pods from the pod template.
//
//	@Summary		Creates the pods from the pod template.
//	@Description	Creates the pods from the pod template, currently the method only supports pod with one container.
//	@Tags			Pod
//	@Accept			application/json
//	@Param			body_param	body	model.UnstructuredCreateRequest	true	"body_param"
//	@Produce		json
//	@Success		201	"Sample result: "{\"message\":\"pod has been created successfully\",\"status\":201}" string
//	@Router			/api/v1/pod/template/create [post]
func createPodFromTemplate(c *fiber.Ctx) error {
	resourceKind := "Pod"
	err := k8s.ParseCreateResource(c, kc, resourceKind)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": resourceKind + " has been created successfully",
	})
}

// createPodFromBody creates the pods from the request body.
//
//	@Summary		Creates the pods from the request body.
//	@Description	Creates the pods from the request body.
//	@Tags			Pod
//	@Accept			application/yaml
//	@Param			body_param	body	string	true	"body_param"
//	@Produce		json
//	@Success		201	"Sample result: "{\"message\":\"pod has been created successfully\",\"status\":201}" string
//	@Router			/api/v1/pod/yaml/create [post]
func createPodFromBody(c *fiber.Ctx) error {
	c.Accepts("application/yaml")
	resourceKind := "Pod"
	if err := k8s.CreateReourceFromData(kc, c.Body(), resourceKind); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error: create " + resourceKind + " failed with error - " + err.Error() + "!",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": resourceKind + " has been created successfully",
	})
}

func deletePod(c *fiber.Ctx) error {
	return nil
}
