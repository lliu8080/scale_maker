package api

import (
	"testing"

	"nuc.lliu.ca/gitea/app/scale_maker/internal/util"
)

func TestPingSuccess(t *testing.T) {
	t.Parallel()
	tests := []util.APITest{
		{
			Description:   "ping route",
			Route:         "/api/v1/ping",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"status\":\"alive\"}",
		},
		{
			Description:   "non existing route",
			Route:         "/api/v1/ping/",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"status\":\"alive\"}",
		},
	}

	app := InitialTestSetup()

	util.RunAPITests(t, app, &tests)
}
