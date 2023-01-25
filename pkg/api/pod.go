package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/form"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

// listPods gets the list of the pods in the k8s cluster.
//
//	@Summary		Gets the list of the pods in the k8s cluster.
//	@Description	Gets the list of the pods in the k8s cluster.
//	@Tags			Kubernetes
//	@Accept			json
//	@Param			namespace	query	string	false	"pod search by namespace"	Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"namespace\":\"default\",\"number_of_pods\":0,\"pods\":[],\"status\":200}" string
//	@Router			/api/v1/pod/list [get]
func listPods(c *fiber.Ctx) error {
	resource := "pods"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "", "v1", resource, namespace)
}

// createPodFromTemplate creates the pods from the pod template.
//
//	@Summary		Creates the pods from the pod template.
//	@Description	Creates the pods from the pod template, currently the method only supports pod with one container.
//	@Tags			Kubernetes
//	@Accept			application/yaml
//	@Param			namespace		path	string	false	"namespace to create the pod"				Format(string)
//	@Param			template_name	path	string	false	"template name needed to create the pod"	Format(string)
//	@Param			cpu_request		path	string	false	"requested cpu value"						Format(string)
//	@Param			memory_request	path	string	false	"requested memory value"					Format(string)
//	@Param			cpu_limit		path	string	false	"limited cpu value"							Format(string)
//	@Param			memory_limit	path	string	false	"limited memory value"						Format(string)
//	@Param			command_params	path	string	false	"Command parameters for pod"				Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"message\":\"pod has been created successfully\",\"status\":200}" string
//	@Router			/api/v1/pod/template/create [post]
func createPodFromTemplate(c *fiber.Ctx) error {
	p := new(form.UnstructuredRequest)
	if err := c.BodyParser(&p); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Error Parsing Request Payload")
	}

	// check template
	cpuLoadTestPodTemplate := "./templates/" + p.TemplateName
	if _, err := util.CheckFileExists(cpuLoadTestPodTemplate); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"message": "Error: unable to retrieve template " +
				cpuLoadTestPodTemplate + " with error " + err.Error() + "!",
		})
	}

	// render template with data
	instanceName := util.GenerateRandomHash()
	data := map[string]string{
		"instanceName":  instanceName,
		"namespace":     p.Namespace,
		"cpuRequest":    p.CPURequest,
		"memoryRequest": p.MemoryRequest,
		"cpuLimit":      p.CPULimit,
		"memoryLimit":   p.MemoryLimit,
		"commandParams": p.CommandParams,
	}

	if err := k8s.CreateReourceFromTempate(kc, cpuLoadTestPodTemplate, data); err != nil {
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
//
//	@Summary		Creates the pods from the request body.
//	@Description	Creates the pods from the request body.
//	@Tags			Kubernetes
//	@Accept			application/yaml
//	@Param			body_param	body	string	true	"body_param"
//	@Produce		json
//	@Success		200	"Sample result: "{\"message\":\"pod has been created successfully\",\"status\":200}" string
//	@Router			/api/v1/pod/yaml/create [post]
func createPodFromBody(c *fiber.Ctx) error {
	c.Accepts("application/yaml")
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
