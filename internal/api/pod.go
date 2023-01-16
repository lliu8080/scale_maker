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

	podList, err := k8s_client.CoreV1().Pods("").List(context.TODO(), meta_v1.ListOptions{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error getting result when trying to list pods!"})
	}
	pods := make([]string, 0, len(podList.Items))
	for _, pod := range podList.Items {
		pods = append(pods, pod.Name)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":         http.StatusOK,
		"number_of_pods": len(podList.Items),
		"pods":           pods,
	})
}
