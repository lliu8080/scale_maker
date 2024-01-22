package k8s

import (
	"bytes"
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/model"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

// ParseCreateResource doc
func ParseCreateResource(c *fiber.Ctx, kc KClient, resourceKind string) error {
	p := new(model.UnstructuredCreateRequest)
	if err := c.BodyParser(&p); err != nil {
		return errors.New("Error parsing request payload with error: " + err.Error())
	}

	// check template
	cpuLoadTestJobTemplate := "./templates/" + p.TemplateName
	if _, err := util.CheckFileExists(cpuLoadTestJobTemplate); err != nil {
		return errors.New("Error: unable to retrieve template " +
			cpuLoadTestJobTemplate + " with error - " + err.Error() + "!")
	}

	// render template with data
	instanceName, err := util.GenerateRandomHash(8)
	if err != nil {
		return errors.New("Error: unable to generate instance name with error - " +
			err.Error() + "!")
	}
	var label string
	if p.TestLabel == "" {
		label = "test"
	} else {
		label = p.TestLabel
	}
	data := map[string]string{
		"testLabel":     label,
		"instanceName":  instanceName,
		"namespace":     p.Namespace,
		"image":         STRESS_TEST_IMAGE,
		"commandParams": p.CommandParams,
		"cpuRequest":    p.CPURequest,
		"memoryRequest": p.MemoryRequest,
		"cpuLimit":      p.CPULimit,
		"memoryLimit":   p.MemoryLimit,
	}

	if err := CreateResourceFromTemplate(
		kc, cpuLoadTestJobTemplate, data, resourceKind); err != nil {
		return errors.New("Error: create " + resourceKind + " failed with error - " +
			err.Error() + "!")
	}
	return nil
}

// CreateResourceFromTemplate - doc
func CreateResourceFromTemplate(kc KClient, templateFullPath string,
	templateData map[string]string, resourceKind string) error {
	resource, err := renderResourceFromTemplate(templateFullPath, templateData)
	if err != nil {
		log.Println("Error: can not load " + templateFullPath + "with error - " + err.Error())
		return err
	}
	return CreateResourceFromData(kc, resource, resourceKind)
}

// CreateResourceFromData - doc
func CreateResourceFromData(kc KClient, data []byte, resourceKind string) error {
	resources, err := serializeResources(data, resourceKind)
	if err != nil {
		log.Println("Error: can not serialize the resource, error - " + err.Error())
		return err
	}
	resourceNum := len(resources)
	if resourceNum == 0 {
		return errors.New("Error: unable to read resource from data")
	}
	for i := 0; i < resourceNum; i++ {
		if err := createResource(kc, resources[i]); err != nil {
			return err
		}
	}
	return nil
}

func serializeResources(data []byte, resourceKind string) ([]model.UnstructuredObj, error) {
	objs := []model.UnstructuredObj{}
	decodedFile := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(data), 100)
	for {
		var rawObj runtime.RawExtension
		var unstructuredObj model.UnstructuredObj
		if err := decodedFile.Decode(&rawObj); err != nil {
			break
		}
		obj, gvk, err := yaml.NewDecodingSerializer(
			unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		if err != nil {
			log.Println("Error: can not serialize data, " + err.Error())
			return []model.UnstructuredObj{}, err
		}
		unstructuredObj.GroupKind = gvk.GroupKind()
		unstructuredObj.Version = gvk.Version
		unstructuredObj.Obj, err = runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			log.Println("Error: can not decode unstructured data, error - " + err.Error())
			return []model.UnstructuredObj{}, err
		}

		if resourceKind != "" && resourceKind != unstructuredObj.Obj["kind"].(string) {
			err := errors.New("Data did not match resource kind " + resourceKind)
			return []model.UnstructuredObj{}, err
		}
		objs = append(objs, unstructuredObj)
	}
	return objs, nil
}

func createResource(kc KClient, obj model.UnstructuredObj) error {

	unstructuredObj := &unstructured.Unstructured{Object: obj.Obj}
	gr, err := restmapper.GetAPIGroupResources(kc.Discovery)
	if err != nil {
		log.Println("Error:  can not get API group resources, " + err.Error())
		return err
	}

	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	mapping, err := mapper.RESTMapping(obj.GroupKind, obj.Version)
	if err != nil {
		log.Println("Error: can not get rest mapping, " + err.Error())
		return err
	}

	var dri dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		if unstructuredObj.GetNamespace() == "" {
			unstructuredObj.SetNamespace("default")
		}
		dri = kc.DynamicClient.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
	} else {
		dri = kc.DynamicClient.Resource(mapping.Resource)
	}

	if _, err := dri.Create(context.Background(), unstructuredObj, metav1.CreateOptions{}); err != nil {
		log.Println("Error: can not create resource, error - " + err.Error())
		return err
	}
	return nil
}
