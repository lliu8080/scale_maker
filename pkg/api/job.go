package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listJobs gets the list of the job in the k8s cluster.
//
//	@Summary		Gets the list of the job in the k8s cluster.
//	@Description	Gets the list of the job in the k8s cluster.
//	@tags			Job
//	@Accept			json
//	@Param			namespace	query	string	false	"job search by namespace"						Format(string)
//	@Param			label		query	string	false	"search job by label"							Format(string)
//	@Param			name		query	string	false	"return job result by name with more details."	Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"jobs\":[],\"namespace\":\"default\",\"number_of_jobs\":0,\"status\":200}" string
//	@Router			/api/v1/job/list [get]
func listJobs(c *fiber.Ctx) error {
	resource := "jobs"
	return k8s.ListResources(c, kc, "batch", "v1", resource)
}

// createJobFromTemplate creates the jobs from the job template.
//
//	@Summary		Creates the jobs from the job template.
//	@Description	Creates the jobs from the job template, currently the method only supports job with one container.
//	@tags			Job
//	@Accept			application/json
//	@Param			body_param	body	model.UnstructuredCreateRequest	true	"body_param"
//	@Produce		json
//	@Success		201	"Sample result: "{\"message\":\"job has been created successfully\",\"status\":201}" string
//	@Router			/api/v1/job/template/create [post]
func createJobFromTemplate(c *fiber.Ctx) error {
	resourceKind := "Job"
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

// createJobFromBody creates the jobs from the request body.
//
//	@Summary		Creates the jobs from the request body.
//	@Description	Creates the jobs from the request body.
//	@tags			Job
//	@Accept			application/yaml
//	@Param			body_param	body	string	true	"body_param"
//	@Produce		json
//	@Success		201	"Sample result: "{\"message\":\"job has been created successfully\",\"status\":201}" string
//	@Router			/api/v1/job/yaml/create [post]
func createJobFromBody(c *fiber.Ctx) error {
	c.Accepts("application/yaml")
	resourceKind := "Job"
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
