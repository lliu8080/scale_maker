package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// listNodes gets the list of the nodes in the k8s cluster.
//
//	@Summary		Gets the list of the nodes in the k8s cluster.
//	@Description	Gets the list of the nodes in the k8s cluster.
//	@Tags			Node
//	@Accept			json
//	@Produce		json
//	@Success		200	"Sample result: "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}""	string
//	@Router			/api/v1/node/list [get]
func listNodes(c *fiber.Ctx) error {
	if kc.ClientSet == nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error initialize k8s client!",
			},
		)
	}

	nList, err := kc.ClientSet.CoreV1().Nodes().List(
		kc.Ctx, metav1.ListOptions{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error getting result when trying to list pods!",
			},
		)
	}
	ns := make(map[string]map[string]string)
	for _, node := range nList.Items {
		nConditions := make(map[string]string)
		//ns[node.Name] = make(map[string]string)
		for _, condition := range node.Status.Conditions {
			nConditions[string(condition.Type)] = string(condition.Status)
		}
		ns[node.Name] = nConditions
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":          http.StatusOK,
		"number_of_nodes": len(nList.Items),
		"nodes":           ns,
	})
}
