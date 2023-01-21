package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listPods gets the list of the pods in the k8s cluster.
// @Summary Gets the list of the pods in the k8s cluster.
// @Description Gets the list of the pods in the k8s cluster.
// @Tags Kubernetes
// @Accept  json
// @Produce  json
// @Success 200 "Sample result: "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}"" string
// @Router /api/v1/pod/list [get]
func listPods(c *fiber.Ctx) error {
	resource := "pods"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "", "v1", resource, namespace)
}

// createPodFromTemplate creates the pods from the pod template.
// @Summary Creates the pods from the pod template.
// @Description Creates the pods from the pod template.
// @Tags Kubernetes
// @Accept  json
// @Produce  json
// @Success 200 "Sample result: "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}"" string
// @Router /api/v1/pod/template/create [post]
func createPodFromTemplate(c *fiber.Ctx) error {
	cpuLoadTestPodTemplate := "./templates/cpu_load_test_pod.yaml"
	if err := k8s.CreateReourceFromTempate(kc, cpuLoadTestPodTemplate); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error: create pod failed with error " + err.Error() + "!",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "pod has been created successfully",
	})
}

// createPodFromBody creates the pods from the request body.
// @Summary Creates the pods from the request body.
// @Description Creates the pods from the request body.
// @Tags Kubernetes
// @Accept  json
// @Produce  json
// @Success 200 "Sample result: "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}"" string
// @Router /api/v1/pod/yaml/create [post]
func createPodFromBody(c *fiber.Ctx) error {
	if err := k8s.CreateReourceFromData(kc, c.Body()); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error: create pod failed with error " + err.Error() + "!",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "pod has been created successfully",
	})
}
