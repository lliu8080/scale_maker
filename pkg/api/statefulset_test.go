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

func setupStatefulsets(statefulsetNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	statefulsets := []runtime.Object{}
	if statefulsetNum != 0 {
		for i := 0; i < statefulsetNum; i++ {
			statefulsets = append(statefulsets, k8s.NewUnstructured("apps/v1", "statefulset", "default", fmt.Sprintf("test-statefulset-%d", i)))
		}
	}
	kc.DynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "apps", Version: "v1", Resource: "statefulsets"}: "statefulsetList",
		},
		statefulsets...,
	)
}

func TestListEmptyStatefulsetSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list statefulsets",
			Route:         "/api/v1/statefulset/list",
			HTTPMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_statefulsets\":0,\"statefulsets\":[],\"status\":200}",
		},
	}
	setupStatefulsets(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMutiStatefulsetSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list statefulsets",
			Route:         "/api/v1/statefulset/list",
			HTTPMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_statefulsets\":2,\"statefulsets\":[\"test-statefulset-0\",\"test-statefulset-1\"],\"status\":200}",
		},
	}
	setupStatefulsets(2)
	util.RunAPITests(t, testApp, &tests)
}
