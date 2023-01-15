package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListNamespace
func ListNamespace(c *fiber.Ctx) error {
	k8s_client := GetLocal[kubernetes.Interface](c, "k8s_client")

	nsList, _ := k8s_client.CoreV1().Namespaces().List(context.Background(), meta_v1.ListOptions{})

	return c.Status(200).JSON(fiber.Map{
		"namespaces": nsList,
	})
}
