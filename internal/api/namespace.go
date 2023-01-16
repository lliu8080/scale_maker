package api

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListNamespace
func ListNamespace(c *fiber.Ctx) error {
	if k8s_client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error initialize k8s client!"})
	}

	nsList, err := k8s_client.CoreV1().Namespaces().List(context.Background(), meta_v1.ListOptions{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error getting result when trying to list pods!"})
	}
	nss := make([]string, 0, len(nsList.Items))
	for _, ns := range nsList.Items {
		nss = append(nss, ns.Name)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":         http.StatusOK,
		"number_of_pods": len(nsList.Items),
		"namespaces":     nss,
	})
}
