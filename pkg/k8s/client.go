package k8s

import (
	"context"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// KClient - document
type KClient struct {
	ClientSet     kubernetes.Interface //*kubernetes.Clientset or fake
	DynamicClient dynamic.Interface    //*dynamic.DynamicClient or fake
	Ctx           context.Context
}
