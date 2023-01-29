package k8s

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ListResources doc
func ListResources(c *fiber.Ctx, kc KClient, group, version,
	resource string) error {

	var err error
	namespace := c.Query("namespace")
	label := c.Query("label")
	workloadName := c.Query("name")
	if namespace == "" {
		namespace = "default"
	}

	listOption := metav1.ListOptions{
		LabelSelector: label,
	}

	if kc.DynamicClient == nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error initialize k8s client!",
			},
		)
	}

	// TODO: Investigate better ways to handle workloadName and list names
	if workloadName != "" {
		listOption := metav1.ListOptions{
			LabelSelector: label,
			FieldSelector: "metadata.name=" + workloadName,
		}
		itemList, err := listDynamicK8SObjectByItems(
			kc,
			listOption,
			group,
			version,
			resource,
			namespace,
		)
		if err != nil || len(itemList) == 0 {
			log.Println("Error: failed list " + resource +
				" with error " + err.Error())
			return c.Status(http.StatusInternalServerError).JSON(
				fiber.Map{
					"status":  http.StatusInternalServerError,
					"message": "Error getting result when trying to list " + resource + "!",
				},
			)
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":                http.StatusOK,
			"namespace":             namespace,
			"number_of_" + resource: len(itemList),
			resource:                itemList,
		})
	}

	stringList, err := listDynamicK8SObjectNames(
		kc,
		listOption,
		group,
		version,
		resource,
		namespace,
	)

	if err != nil {
		log.Println("Error: failed list " + resource +
			" with error " + err.Error())
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": "Error getting result when trying to list " + resource + "!",
			},
		)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":                http.StatusOK,
		"namespace":             namespace,
		"number_of_" + resource: len(stringList),
		resource:                stringList,
	})
}

// listDynamicK8SObjectByItems doc
func listDynamicK8SObjectByItems(kc KClient,
	listOption metav1.ListOptions, group, version, resource, namespace string) (
	[]unstructured.Unstructured, error) {
	resourceID := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	list, err := kc.DynamicClient.Resource(resourceID).Namespace(namespace).
		List(kc.Ctx, listOption)
	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

// listDynamicK8SObjectByNames doc
func listDynamicK8SObjectNames(kc KClient,
	listOption metav1.ListOptions, group, version, resource, namespace string) (
	[]string, error) {

	items, err := listDynamicK8SObjectByItems(kc, listOption,
		group, version, resource, namespace)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0, len(items))
	for _, item := range items {
		results = append(results, item.GetName())
	}

	return results, nil
}
