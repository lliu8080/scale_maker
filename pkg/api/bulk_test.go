package api

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func setupResources() {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	resources := []runtime.Object{}
	kc.DynamicClient = fake.NewSimpleDynamicClient(scheme, resources...)
}

func TestCreateResourcesFromBodySuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:         "Create pods from request body",
			Route:               "/api/v1/bulk/create",
			HTTPMethod:          "POST",
			RequestBodyFromFile: "../../test_data/pkg/api/bulk/test_bulk.yaml",
			ExpectedError:       false,
			ExpectedCode:        201,
			ExpectedBody:        "{\"message\":\"k8s resources have been created successfully\",\"status\":201}",
		},
	}
	setupResources()
	util.RunAPITests(t, testApp, &tests)
}
