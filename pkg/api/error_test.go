package api

import (
	"testing"

	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func TestNotFound(t *testing.T) {
	tests := []util.APITest{
		{
			Description:   "ping invalid url",
			Route:         "/api/v1/invalid_url",
			HTTPMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  404,
			ExpectedBody:  "{\"status\":\"Error: 404 page not found!\"}",
		},
	}
	app := InitialTestSetup()
	util.RunAPITests(t, app, &tests)
}
