package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// listNamespace doc
func listNamespaces(c *fiber.Ctx) error {
	if kc.clientSet == nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error initialize k8s client!",
			},
		)
	}

	nsList, err := kc.clientSet.CoreV1().Namespaces().List(
		kc.ctx, metav1.ListOptions{})
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
