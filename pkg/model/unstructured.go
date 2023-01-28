package model

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UnstructuredObj doc
type UnstructuredObj struct {
	Obj       map[string]interface{}
	GroupKind schema.GroupKind
	Version   string
}

// UnstructuredCreateRequest doc
type UnstructuredCreateRequest struct {
	Namespace     string `json:"namespace" validate:"required"`
	TemplateName  string `json:"templateName" validate:"required,min=2,max=50"`
	CPURequest    string `json:"cpuRequest"`
	MemoryRequest string `json:"memoryRequest"`
	CPULimit      string `json:"cpuLimit"`
	MemoryLimit   string `json:"memoryLimit"`
	CommandParams string `json:"commandParams" validate:"required"`
	TestLabel     string `json:"testLabel"`
}

// UnstructuredDeleteRequest doc
type UnstructuredDeleteRequest struct {
	TestLabel string `json:"test_label" validate:"required"`
}
