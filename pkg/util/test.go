package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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
			// tests need to load test files
			if strings.Contains(test.RequestBodyFromFile, ".yaml") {
				// process yaml test files
				testData, err = ioutil.ReadFile(test.RequestBodyFromFile)
				data, err = yaml.YAMLToJSONStrict(testData)
				header = "application/yaml"
			} else {
				// process json test files
				testData, err = ioutil.ReadFile(test.RequestBodyFromFile)
				err = json.Unmarshal(testData, &tmp)
				data, err = json.Marshal(tmp)
				header = "application/json"
			}
		} else {
			// tests don't need to load test files
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

		// verify that no error occurred, that is not expected
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

		// Reading the response body should work every time, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.Description)

		// Verify, that the response body equals the expected body
		assert.Equalf(t, test.ExpectedBody, string(body), test.Description)
	}
}
