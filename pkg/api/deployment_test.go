package api

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func setupDeployments(deploymentNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	deployments := []runtime.Object{}
	if deploymentNum != 0 {
		for i := 0; i < deploymentNum; i++ {
			deployments = append(deployments, newUnstructured("apps/v1", "deployment", "default", fmt.Sprintf("test-deployment-%d", i)))
		}
	}
	kc.dynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "apps", Version: "v1", Resource: "deployments"}: "deploymentList",
		},
		deployments...,
	)
}

func TestListEmptyDeploymentSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list deployments",
			Route:         "/api/v1/deployment/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"deployments\":[],\"namespace\":\"default\",\"number_of_deployments\":0,\"status\":200}",
		},
	}
	setupDeployments(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMutiDeploymentSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list deployments",
			Route:         "/api/v1/deployment/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"deployments\":[\"test-deployment-0\",\"test-deployment-1\"],\"namespace\":\"default\",\"number_of_deployments\":2,\"status\":200}",
		},
	}
	setupDeployments(2)
	util.RunAPITests(t, testApp, &tests)
}
