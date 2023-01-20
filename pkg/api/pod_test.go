package api

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func newUnstructured(apiVersion, kind, namespace, name string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"namespace": namespace,
				"name":      name,
			},
		},
	}
}

func setupPods(podNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	pods := []runtime.Object{}
	if podNum != 0 {
		for i := 0; i < podNum; i++ {
			pods = append(pods, newUnstructured("v1", "pod", "default", fmt.Sprintf("test-pod-%d", i)))
		}
	}
	kc.dynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "", Version: "v1", Resource: "pods"}: "podList",
		},
		pods...,
	)

}

func TestListEmptyPodSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list namespaces",
			Route:         "/api/v1/pod/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_pods\":0,\"pods\":[],\"status\":200}",
		},
	}
	setupPods(0)
	util.RunAPITests(t, testApp, &tests)
}
