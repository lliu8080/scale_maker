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
        "/api/v1/bulk/create": {
            "post": {
                "description": "Creates all the resources passed via request body.",
                "consumes": [
                    "application/yaml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes Bulk API"
                ],
                "summary": "Creates all the resources passed via request body.",
                "parameters": [
                    {
                        "description": "body_param",
                        "name": "body_param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"daemonsets\\\":[],\\\"namespace\\\":\\\"default\\\",\\\"number_of_daemonsets\\\":0,\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/daemonset/list": {
            "get": {
                "description": "Gets the list of the daemonsets in the k8s cluster.",
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
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "daemonset search by namespace",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"daemonsets\\\":[],\\\"namespace\\\":\\\"default\\\",\\\"number_of_daemonsets\\\":0,\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/deployment/list": {
            "get": {
                "description": "Gets the list of the deployments in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the deployments in the k8s cluster.",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "deployment search by namespace",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"deployments\\\":[],\\\"namespace\\\":\\\"default\\\",\\\"number_of_deployments\\\":0,\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/namespace/list": {
            "get": {
                "description": "Gets the list of the namespaces in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the namespaces in the k8s cluster.",
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"namespaces\\\":[],\\\"number_of_namespaces\\\":0,\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/node/list": {
            "get": {
                "description": "Gets the list of the nodes in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the nodes in the k8s cluster.",
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
        },
        "/api/v1/pod/list": {
            "get": {
                "description": "Gets the list of the pods in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the pods in the k8s cluster.",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "pod search by namespace",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"namespace\\\":\\\"default\\\",\\\"number_of_pods\\\":0,\\\"pods\\\":[],\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/pod/template/create": {
            "post": {
                "description": "Creates the pods from the pod template, currently the method only supports pod with one container.",
                "consumes": [
                    "application/yaml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Creates the pods from the pod template.",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "namespace to create the pod",
                        "name": "namespace",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "template name needed to create the pod",
                        "name": "template_name",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "requested cpu value",
                        "name": "cpu_request",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "requested memory value",
                        "name": "memory_request",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "limited cpu value",
                        "name": "cpu_limit",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "limited memory value",
                        "name": "memory_limit",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "Command parameters for pod",
                        "name": "command_params",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"message\\\":\\\"pod has been created successfully\\\",\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/pod/yaml/create": {
            "post": {
                "description": "Creates the pods from the request body.",
                "consumes": [
                    "application/yaml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Creates the pods from the request body.",
                "parameters": [
                    {
                        "description": "body_param",
                        "name": "body_param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"message\\\":\\\"pod has been created successfully\\\",\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/service/list": {
            "get": {
                "description": "Gets the list of the services in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the services in the k8s cluster.",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "service search by namespace",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"namespace\\\":\\\"default\\\",\\\"number_of_services\\\":0,\\\"services\\\":[],\\\"status\\\":200}"
                    }
                }
            }
        },
        "/api/v1/statefulset/list": {
            "get": {
                "description": "Gets the list of the statefulsets in the k8s cluster.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Gets the list of the statefulsets in the k8s cluster.",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "statefulset search by namespace",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sample result: \"{\\\"namespace\\\":\\\"default\\\",\\\"number_of_statefulsets\\\":0,\\\"statefulsets\\\":[],\\\"status\\\":200}"
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
