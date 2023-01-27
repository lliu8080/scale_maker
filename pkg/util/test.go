package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/k8s"
	"sigs.k8s.io/yaml"
)

// APITest doc
type APITest struct {
	Description string

	// Test input
	Route               string
	HTTPMethod          string
	RequestBodyFromFile string

	// Expected output
	ExpectedError bool
	ExpectedCode  int
	ExpectedBody  string
}

// RunAPITests doc
func RunAPITests(t *testing.T, app *fiber.App, tests *[]APITest) {

	// Iterate through test single test cases
	for _, test := range *tests {
		// Create a new http request with the route
		// from the test case
		var (
			testData []byte
			tmp      interface{}
			data     []byte
			err      error
			header   string
		)
		if test.RequestBodyFromFile != "" {
			if strings.Contains(test.RequestBodyFromFile, ".yaml") {
				testData, err = ioutil.ReadFile(test.RequestBodyFromFile)
				data, err = yaml.YAMLToJSONStrict(testData)
				header = "application/yaml"
			} else {
				testData, err = ioutil.ReadFile(test.RequestBodyFromFile)
				err = json.Unmarshal(testData, &tmp)
				data, err = json.Marshal(tmp)
				header = "application/json"
			}
		} else {
			data, _ = json.Marshal(test.RequestBodyFromFile)
			header = "application/json"
		}

		req, _ := http.NewRequest(
			test.HTTPMethod,
			test.Route,
			bytes.NewBuffer(data),
		)
		req.Header.Set("Content-type", header)
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.ExpectedError, err != nil, test.Description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.ExpectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.ExpectedCode, res.StatusCode, test.Description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.Description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.ExpectedBody, string(body), test.Description)
	}
}

// SetupDiscovery doc
func SetupDiscovery(kc k8s.KClient) discovery.DiscoveryInterface {
	fakeDiscovery, ok := kc.ClientSet.Discovery().(*fakediscovery.FakeDiscovery)
	if !ok {
		log.Fatalf("couldn't convert Discovery() to *FakeDiscovery")
	}
	fakeDiscovery.Fake.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: "v1",
			APIResources: []metav1.APIResource{
				{
					Kind: "Pod",
					Name: "Pods",
				},
				{
					Kind: "Service",
					Name: "Services",
				},
			},
		},
		{
			GroupVersion: "batch/v1",
			APIResources: []metav1.APIResource{
				{
					Kind: "Job",
					Name: "Jobs",
				},
			},
		},
		{
			GroupVersion: "apps/v1",
			APIResources: []metav1.APIResource{
				{
					Kind: "Deployment",
					Name: "Deployments",
				},
			},
		},
	}
	return fakeDiscovery
}
