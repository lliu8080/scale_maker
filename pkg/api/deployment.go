package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listDeployments gets the list of the deployments in the k8s cluster.
//
//	@Summary		Gets the list of the deployments in the k8s cluster.
//	@Description	Gets the list of the deployments in the k8s cluster.
//	@Tags			Deployments
//	@Accept			json
//	@Param			namespace	query	string	false	"deployment search by namespace"															Format(string)
//	@Param			label		query	string	false	"search deployment by label"																Format(string)
//	@Param			by_item		query	string	false	"set by_item=true to return deployment results by item with more details, default false."	Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"deployments\":[],\"namespace\":\"default\",\"number_of_deployments\":0,\"status\":200}"	string
//	@Router			/api/v1/deployment/list [get]
func listDeployments(c *fiber.Ctx) error {
	resource := "deployments"
	return k8s.ListResources(c, kc, "apps", "v1", resource)
}
