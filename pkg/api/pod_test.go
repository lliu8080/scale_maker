package api

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func setupPods(podNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	pods := []runtime.Object{}
	if podNum != 0 {
		for i := 0; i < podNum; i++ {
			pods = append(pods, k8s.NewUnstructured("v1", "Pod", "default", fmt.Sprintf("test-pod-%d", i)))
		}
	}
	//scheme.AddKnownTypeWithName(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "pods"}, &ruPod{})
	kc.DynamicClient = fake.NewSimpleDynamicClient(scheme, pods...)
}

func TestListEmptyPodSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list pods",
			Route:         "/api/v1/pod/list",
			HTTPMethod:    "GET",
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
			HTTPMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespace\":\"default\",\"number_of_pods\":2,\"pods\":[\"test-pod-0\",\"test-pod-1\"],\"status\":200}",
		},
	}
	setupPods(2)
	util.RunAPITests(t, testApp, &tests)
}

func TestCreatePodFromTemplateSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:         "Create pods from request body",
			Route:               "/api/v1/pod/template/create",
			HTTPMethod:          "POST",
			RequestBodyFromFile: "../../test_data/pkg/api/pod/test_params.json",
			ExpectedError:       false,
			ExpectedCode:        201,
			ExpectedBody:        "{\"message\":\"Pod has been created successfully\",\"status\":201}",
		},
	}
	setupPods(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestCreatePodFromBodySuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:         "Create pods from request body",
			Route:               "/api/v1/pod/yaml/create",
			HTTPMethod:          "POST",
			RequestBodyFromFile: "../../test_data/pkg/api/pod/test_pod.yaml",
			ExpectedError:       false,
			ExpectedCode:        201,
			ExpectedBody:        "{\"message\":\"Pod has been created successfully\",\"status\":201}",
		},
	}
	setupPods(0)
	util.RunAPITests(t, testApp, &tests)
}
