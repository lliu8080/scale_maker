package api

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListPod
func ListPod(c *fiber.Ctx) error {
	if k8s_client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error initialize k8s client!"})
	}

	nsList, err := k8s_client.CoreV1().Namespaces().List(context.TODO(), meta_v1.ListOptions{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error getting result when trying to list pods!"})
	}
	return c.Status(200).JSON(fiber.Map{
		"pods": nsList,
	})
}
