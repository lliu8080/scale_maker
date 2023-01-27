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

func setupJobs(jobNum int) {
	testApp = InitialTestSetup()
	scheme := runtime.NewScheme()
	jobs := []runtime.Object{}
	if jobNum != 0 {
		for i := 0; i < jobNum; i++ {
			jobs = append(jobs, k8s.NewUnstructured("batch/v1", "Job", "default", fmt.Sprintf("test-job-%d", i)))
		}
	}
	kc.DynamicClient = fake.NewSimpleDynamicClientWithCustomListKinds(scheme,
		map[schema.GroupVersionResource]string{
			{Group: "batch", Version: "v1", Resource: "jobs"}: "jobList",
		},
		jobs...,
	)
}

func TestListEmptyJobSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list jobs",
			Route:         "/api/v1/job/list",
			HTTPMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"jobs\":[],\"namespace\":\"default\",\"number_of_jobs\":0,\"status\":200}",
		},
	}
	setupJobs(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMutiJobSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "list jobs",
			Route:         "/api/v1/job/list",
			HTTPMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"jobs\":[\"test-job-0\",\"test-job-1\"],\"namespace\":\"default\",\"number_of_jobs\":2,\"status\":200}",
		},
	}
	setupJobs(2)
	util.RunAPITests(t, testApp, &tests)
}

func TestCreateJobFromTemplateSuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:         "Create pods from request body",
			Route:               "/api/v1/job/template/create",
			HTTPMethod:          "POST",
			RequestBodyFromFile: "../../test_data/pkg/api/job/test_params.json",
			ExpectedError:       false,
			ExpectedCode:        201,
			ExpectedBody:        "{\"message\":\"Job has been created successfully\",\"status\":201}",
		},
	}
	setupJobs(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestCreateJobFromBodySuccess(t *testing.T) {
	tests := []util.APITest{
		{
			Description:         "Create pods from request body",
			Route:               "/api/v1/job/yaml/create",
			HTTPMethod:          "POST",
			RequestBodyFromFile: "../../test_data/pkg/api/job/test_job.yaml",
			ExpectedError:       false,
			ExpectedCode:        201,
			ExpectedBody:        "{\"message\":\"Job has been created successfully\",\"status\":201}",
		},
	}
	setupJobs(0)
	util.RunAPITests(t, testApp, &tests)
}
