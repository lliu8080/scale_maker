package k8s

import (
	"bytes"
	"text/template"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// NewUnstructured - doc
func NewUnstructured(apiVersion, kind, namespace, name string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"namespace": namespace,
				"name":      name,
			},
		},
	}
}

// renderResourceFromTemplate doc
func renderResourceFromTemplate(templatePath string, data map[string]string) ([]byte, error) {
	var tpl bytes.Buffer

	t, err := template.New("new").ParseFiles(templatePath)
	if err != nil {
		return []byte{}, err
	}

	if err := t.Execute(&tpl, data); err != nil {
		return []byte{}, err
	}
	return tpl.Bytes(), nil
}
