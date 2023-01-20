package api

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func setupDaemonsets(daemonsetNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	daemonsets := []runtime.Object{}
	if daemonsetNum != 0 {
		for i := 0; i < daemonsetNum; i++ {
			daemonsets = append(daemonsets, newUnstructured("apps/v1", "daemonset", "default", fmt.Sprintf("test-daemonset-%d", i)))
		}
	}
	kc.dynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "apps", Version: "v1", Resource: "daemonsets"}: "daemonsetList",
		},
		daemonsets...,
	)
}

func TestListEmptyDaemonsetSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list daemonsets",
			Route:         "/api/v1/daemonset/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"daemonsets\":[],\"namespace\":\"default\",\"number_of_daemonsets\":0,\"status\":200}",
		},
	}
	setupDaemonsets(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMutiDaemonsetSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list daemonsets",
			Route:         "/api/v1/daemonset/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"daemonsets\":[\"test-daemonset-0\",\"test-daemonset-1\"],\"namespace\":\"default\",\"number_of_daemonsets\":2,\"status\":200}",
		},
	}
	setupDaemonsets(2)
	util.RunAPITests(t, testApp, &tests)
}
