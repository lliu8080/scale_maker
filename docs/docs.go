// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/daemonset/list": {
            "get": {
                "description": "Gets the list of the Daemonsets in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the Daemonsets in the k8s cluster.",
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"daemonsets\\\":[],\\\"namespace\\\":\\\"default\\\",\\\"number_of_daemonsets\\\":0,\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/ping": {
            "get": {
                "description": "Fetch the current status of the application.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Fetch the current status of the application.",
                "responses": {
                    "200": {
                        "description": "Sample result: {\\\"status\\\":\\\"alive\\\"}"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
