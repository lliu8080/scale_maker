package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// listNamespaces gets the list of the namespaces in the k8s cluster.
//
//	@Summary		Gets the list of the namespaces in the k8s cluster.
//	@Description	Gets the list of the namespaces in the k8s cluster.
//	@Tags			Namespace
//	@Accept			json
//	@Produce		json
//	@Success		200	"Sample result: "{\"namespaces\":[],\"number_of_namespaces\":0,\"status\":200}"	string
//	@Router			/api/v1/namespace/list [get]
func listNamespaces(c *fiber.Ctx) error {
	if kc.ClientSet == nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error initialize k8s client!",
			},
		)
	}

	nsList, err := kc.ClientSet.CoreV1().Namespaces().List(
		kc.Ctx, metav1.ListOptions{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error getting result when trying to list pods!",
			},
		)
	}

	nss := make([]string, 0, len(nsList.Items))
	for _, ns := range nsList.Items {
		nss = append(nss, ns.Name)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":               http.StatusOK,
		"number_of_namespaces": len(nsList.Items),
		"namespaces":           nss,
	})
}
