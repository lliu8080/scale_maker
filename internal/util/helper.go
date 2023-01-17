package util

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// func getK8SObjectNames[T any](objs []T) []T {
// 	res := make([]T, 0, len(objs))
// 	for _, obj := range objs {
// 		res = append(res, obj.GetType().Name)
// 	}
// 	return res
// }

// ListDynamicK8SObjectByItems doc
func ListDynamicK8SObjectByItems(ctx context.Context, dynamic dynamic.Interface,
	group string, version string, resource string, namespace string) (
	[]unstructured.Unstructured, error) {
	resourceID := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}
	list, err := dynamic.Resource(resourceID).Namespace(namespace).
		List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

// ListDynamicK8SObjectByNames doc
func ListDynamicK8SObjectByNames(ctx context.Context, dynamic dynamic.Interface,
	group string, version string, resource string, namespace string) (
	[]string, error) {

	items, err := ListDynamicK8SObjectByItems(ctx, dynamic, group, version, resource, namespace)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0, len(items))
	for _, item := range items {
		results = append(results, item.GetName())
	}
	return results, nil
}
