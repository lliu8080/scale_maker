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

func setupNode(nodeNum int) {
	testApp = InitialTestSetup()
	ns := []runtime.Object{}
	if nodeNum != 0 {
		for i := 0; i < nodeNum; i++ {
			ns = append(ns, &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name:        fmt.Sprintf("test-node-%d", i),
					Annotations: map[string]string{},
				},
			})
		}
	}
	kc.clientSet = fake.NewSimpleClientset(ns...)
}

func TestListEmptyNodeSuccess(t *testing.T) {
	//t.Parallel()
	tests := []util.APITest{
		{
			Description:   "list nodes",
			Route:         "/api/v1/node/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"nodes\":{},\"number_of_nodes\":0,\"status\":200}",
		},
	}
	setupNode(0)
	util.RunAPITests(t, testApp, &tests)
}

func TestListMultiNodesSuccess(t *testing.T) {
	//t.Parallel()
	tests := []util.APITest{
		{
			Description:   "list nodes",
			Route:         "/api/v1/node/list",
			HttpMethod:    "GET",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "{\"nodes\":{\"test-node-0\":{},\"test-node-1\":{}},\"number_of_nodes\":2,\"status\":200}",
		},
	}
	setupNode(2)
	util.RunAPITests(t, testApp, &tests)
}
