package k8s

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
)

// CreateReourceFromTempate - doc
func CreateReourceFromTempate(kc KClient, templateFullPath string) error {
	cpuLoadTestPodData, err := ioutil.ReadFile(templateFullPath)
	if err != nil {
		log.Println("Error: can not load " + templateFullPath + "with error " + err.Error())
		return err
	}
	return CreateReourceFromData(kc, cpuLoadTestPodData)
}

// CreateReourceFromData - doc
func CreateReourceFromData(kc KClient, data []byte) error {
	var rawObj runtime.RawExtension
	decodedFile := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(data), 100)

	if err := decodedFile.Decode(&rawObj); err != nil {
		log.Println("Error: can not decode data, " + err.Error())
		return err
	}

	obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
	unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		log.Println("Error: can not decode data, " + err.Error())
		return err
	}

	unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}
	gr, err := restmapper.GetAPIGroupResources(kc.ClientSet.Discovery())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	mapper := restmapper.NewDiscoveryRESTMapper(gr)
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		log.Println(err.Error())
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
		log.Println(err.Error())
	}
	return nil
}
