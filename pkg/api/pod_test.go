package api

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func setupPods(podNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	pods := []runtime.Object{}
	if podNum != 0 {
		for i := 0; i < podNum; i++ {
			pods = append(pods, k8s.NewUnstructured("v1", "pod", "default", fmt.Sprintf("test-pod-%d", i)))
		}
	}
	kc.DynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "", Version: "v1", Resource: "pods"}: "podList",
		},
		pods...,
	)
}

func TestListEmptyPodSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list pods",
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

func TestListMutiPodSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list pods",
			Route:         "/api/v1/pod/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_pods\":2,\"pods\":[\"test-pod-0\",\"test-pod-1\"],\"status\":200}",
		},
	}
	setupPods(2)
	util.RunAPITests(t, testApp, &tests)
}
