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
