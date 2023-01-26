package k8s

import (
	"bytes"
	"context"
	"errors"
	"log"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/model"
)

// CreateReourceFromTempate - doc
func CreateReourceFromTempate(
	kc KClient, templateFullPath string, templateData map[string]string, resourceKind string) error {
	resource, err := renderResourceFromTemplate(templateFullPath, templateData)
	if err != nil {
		log.Println("Error: can not load " + templateFullPath + "with error " + err.Error())
		return err
	}
	return CreateReourceFromData(kc, resource, resourceKind)
}

// CreateReourceFromData - doc
func CreateReourceFromData(kc KClient, data []byte, resourceKind string) error {
	resources, err := serializeResources(data, resourceKind)
	if err != nil {
		log.Println("Error: can not serialize the resource, " + err.Error())
		return err
	}
	resourceNum := len(resources)
	if resourceNum == 0 {
		return errors.New("Error: unable to read resource from data")
	}
	for i := 0; i < resourceNum; i++ {
		// if err := validateReource(resources[i]); err != nil {
		// 	return err
		// }
		if err := createReource(kc, resources[i]); err != nil {
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
			log.Println("Error: can not decode unstructured data, " + err.Error())
			return []model.UnstructuredObj{}, err
		}
		if resourceKind != "" && resourceKind != unstructuredObj.Obj["Kind"] {
			err := errors.New("Data did not match resource kind" + resourceKind)
			return []model.UnstructuredObj{}, err
		}
		objs = append(objs, unstructuredObj)
	}
	return objs, nil
}

// func validateReource(obj model.UnstructuredObj) error {
// 	return nil
// }

func createReource(kc KClient, obj model.UnstructuredObj) error {

	unstructuredObj := &unstructured.Unstructured{Object: obj.Obj}
	gr, err := restmapper.GetAPIGroupResources(kc.ClientSet.Discovery())
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
		log.Println("Error: can not create resource, " + err.Error())
		return err
	}
	return nil
}
