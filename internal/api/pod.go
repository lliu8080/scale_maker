package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
)

// listPods doc
func listPods(c *fiber.Ctx) error {
	resource := "pods"
	namespace := c.Query("namespace")
	return listResources(c, "", "v1", resource, namespace)
}

func createPod(c *fiber.Ctx) error {
	var rawObj runtime.RawExtension
	var err error
	if err = kt.cpuLoadTestPod.Decode(&rawObj); err != nil {
		log.Println(err.Error())
	}

	obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
	unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		log.Println(err.Error())
	}

	unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

	gr, err := restmapper.GetAPIGroupResources(kc.clientSet.Discovery())
	if err != nil {
		log.Println(err.Error())
	}

	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		log.Println(err.Error())
	}

	var dri dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		if unstructuredObj.GetNamespace() == "" {
			unstructuredObj.SetNamespace("default")
		}
		dri = kc.dynamicClient.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
	} else {
		dri = kc.dynamicClient.Resource(mapping.Resource)
	}

	if _, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{}); err != nil {
		log.Println(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "pod created successfully",
	})
}
