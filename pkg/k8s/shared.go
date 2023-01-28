package k8s

import (
	"bytes"
	"log"
	"text/template"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
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

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return []byte{}, err
	}

	if err := t.Execute(&tpl, data); err != nil {
		return []byte{}, err
	}
	return tpl.Bytes(), nil
}

// SetupDiscovery doc
func SetupDiscovery(kc KClient) discovery.DiscoveryInterface {
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
