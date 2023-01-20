package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type APITest struct {
	Description string

	// Test input
	Route       string
	HttpMethod  string
	RequestBody map[string]interface{}

	// Expected output
	ExpectedError bool
	ExpectedCode  int
	ExpectedBody  string
}

func RunAPITests(t *testing.T, app *fiber.App, tests *[]APITest) {

	// Iterate through test single test cases
	for _, test := range *tests {
		// Create a new http request with the route
		// from the test case

		data, _ := json.Marshal(test.RequestBody)
		//data := bytes.NewBuffer(req)

		req, _ := http.NewRequest(
			test.HttpMethod,
			test.Route,
			bytes.NewBuffer(data),
		)
		req.Header.Set("Content-type", "application/json")
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
