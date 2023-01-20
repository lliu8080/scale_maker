package api

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"nuc.lliu.ca/gitea/app/scale_maker/pkg/util"
)

func setupNamespace(namespaceNum int) {
	testApp = InitialTestSetup()
	ns := []runtime.Object{}
	if namespaceNum != 0 {
		for i := 0; i < namespaceNum; i++ {
			ns = append(ns, &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        fmt.Sprintf("test-namespace-%d", i),
					Annotations: map[string]string{},
				},
			})
		}
	}
	kc.ClientSet = fake.NewSimpleClientset(ns...)
}

func TestListEmptyNamespaceSuccess(t *testing.T) {
	//t.Parallel()
	tests := []util.APITest{
		{
			Description:   "list namespaces",
			Route:         "/api/v1/namespace/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespaces\":[],\"number_of_namespaces\":0,\"status\":200}",
		},
	}
	setupNamespace(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMultiNamespacesSuccess(t *testing.T) {
	//t.Parallel()
	tests := []util.APITest{
		{
			Description:   "list namespaces",
			Route:         "/api/v1/namespace/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"namespaces\":[\"test-namespace-0\",\"test-namespace-1\"],\"number_of_namespaces\":2,\"status\":200}",
		},
	}
	setupNamespace(2)
	util.RunAPITests(t, testApp, &tests)
}
