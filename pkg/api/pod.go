package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listPods doc
func listPods(c *fiber.Ctx) error {
	resource := "pods"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "", "v1", resource, namespace)
}

// createPod creates a new pod from yaml template
func createPod(c *fiber.Ctx) error {
	cpuLoadTestPodTemplate := "./templates/cpu_load_test_pod.yaml"
	if err := k8s.CreateReourceFromTempate(kc, cpuLoadTestPodTemplate); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Error: create pod failed with error " + err.Error() + "!",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "pod created successfully",
	})
}
