package api

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListNamespace
func ListNamespace(c *fiber.Ctx) error {
	k8s_client := GetLocal[kubernetes.Interface](c, "k8s_client")
	if k8s_client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error initialize k8s client!"})
	}

	nsList, err := k8s_client.CoreV1().Namespaces().List(context.Background(), meta_v1.ListOptions{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error getting result when trying to list pods!"})
	}
	return c.Status(200).JSON(fiber.Map{
		"namespaces": nsList,
	})
}
