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

func setupServices(serviceNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	services := []runtime.Object{}
	if serviceNum != 0 {
		for i := 0; i < serviceNum; i++ {
			services = append(services, k8s.NewUnstructured("core/v1", "service", "default", fmt.Sprintf("test-service-%d", i)))
		}
	}
	kc.DynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "core", Version: "v1", Resource: "services"}: "serviceList",
		},
		services...,
	)
}

func TestListEmptyServiceSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list services",
			Route:         "/api/v1/service/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_services\":0,\"services\":[],\"status\":200}",
		},
	}
	setupServices(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMutiServiceSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list services",
			Route:         "/api/v1/service/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_services\":2,\"services\":[\"test-service-0\",\"test-service-1\"],\"status\":200}",
		},
	}
	setupServices(2)
	util.RunAPITests(t, testApp, &tests)
}
