package api

import (
	"github.com/gofiber/fiber/v2"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
)

// listStatefulsets gets the list of the statefulsets in the k8s cluster.
//
//	@Summary		Gets the list of the statefulsets in the k8s cluster.
//	@Description	Gets the list of the statefulsets in the k8s cluster.
//	@Tags			Statefulsets
//	@Accept			json
//	@Param			namespace	query	string	false	"statefulset search by namespace"						Format(string)
//	@Param			label		query	string	false	"search statefulset by label"							Format(string)
//	@Param			name		query	string	false	"return statefulset result by name with more details."	Format(string)
//	@Produce		json
//	@Success		200	"Sample result: "{\"namespace\":\"default\",\"number_of_statefulsets\":0,\"statefulsets\":[],\"status\":200}"	string
//	@Router			/api/v1/statefulset/list [get]
func listStatefulsets(c *fiber.Ctx) error {
	resource := "statefulsets"
	return k8s.ListResources(c, kc, "apps", "v1", resource)
}
