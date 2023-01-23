package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listServices gets the list of the services in the k8s cluster.
//
//	@Summary		Gets the list of the services in the k8s cluster.
//	@Description	Gets the list of the services in the k8s cluster.
//	@Tags			Kubernetes
//	@Accept			json
//	@Param			namespace	query	string	false	"service search by namespace"	Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"namespace\":\"default\",\"number_of_services\":0,\"services\":[],\"status\":200}"	string
//	@Router			/api/v1/service/list [get]
func listServices(c *fiber.Ctx) error {
	resource := "services"
	namespace := c.Query("namespace")
	return k8s.ListResources(c, kc, "core", "v1", resource, namespace)
}
