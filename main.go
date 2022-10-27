package main

import (
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/runtime"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/client"
	"net/http"
)

func main() {

	s := &runtime.Server{
		ClientSetK8s:      client.GetOutClusterClientSetKubernetes(),
		ClientSetV1alpha1: client.GetOutClusterClientSetV1alpha1(),
	}
	h := gen.Handler(s)

	http.ListenAndServe(":3000", h)
}

//package main
//
//import (
//	"context"
//	"fmt"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/runtime/schema"
//
//	"github.com/itchyny/gojq"
//	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/client-go/dynamic"
//	ctrl "sigs.k8s.io/controller-runtime"
//)
//
//func main() {
//	ctx := context.Background()
//	config := ctrl.GetConfigOrDie()
//	dynamic := dynamic.NewForConfigOrDie(config)
//
//	namespace := "default"
//	query := ".metadata.uid == \"f005d9d3-19fc-4e55-a15e-75fc7303efb1\""
//	items, err := GetResourcesByJq(dynamic, ctx, "gateway.networking.k8s.io", "v1alpha2", "httproutes", namespace, query)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		for _, item := range items {
//			fmt.Printf("%+v\n", item)
//		}
//	}
//}
//
//func GetResourcesByJq(dynamic dynamic.Interface, ctx context.Context, group string,
//	version string, resource string, namespace string, jq string) (
//	[]unstructured.Unstructured, error) {
//
//	resources := make([]unstructured.Unstructured, 0)
//
//	query, err := gojq.Parse(jq)
//	if err != nil {
//		return nil, err
//	}
//
//	items, err := GetResourcesDynamically(dynamic, ctx, group, version, resource, namespace)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, item := range items {
//		// Convert object to raw JSON
//		var rawJson interface{}
//		err = runtime.DefaultUnstructuredConverter.FromUnstructured(item.Object, &rawJson)
//		if err != nil {
//			return nil, err
//		}
//
//		// Evaluate jq against JSON
//		iter := query.Run(rawJson)
//		for {
//			result, ok := iter.Next()
//			if !ok {
//				break
//			}
//			if err, ok := result.(error); ok {
//				if err != nil {
//					return nil, err
//				}
//			} else {
//				boolResult, ok := result.(bool)
//				if !ok {
//					fmt.Println("Query returned non-boolean value")
//				} else if boolResult {
//					resources = append(resources, item)
//				}
//			}
//		}
//	}
//	return resources, nil
//}
//
//func GetResourcesDynamically(dynamic dynamic.Interface, ctx context.Context,
//	group string, version string, resource string, namespace string) (
//	[]unstructured.Unstructured, error) {
//
//	resourceId := schema.GroupVersionResource{
//		Group:    group,
//		Version:  version,
//		Resource: resource,
//	}
//	//list, err := dynamic.Resource(resourceId).Namespace(namespace).
//	//	List(ctx, metav1.ListOptions{})
//	list, err := dynamic.Resource(resourceId).Namespace(namespace).
//		List(ctx, metav1.ListOptions{})
//
//	if err != nil {
//		return nil, err
//	}
//
//	return list.Items, nil
//}

//func main() {
//	c := openapi3.T{
//		ExtensionProps: openapi3.ExtensionProps{},
//		OpenAPI:        "test",
//		Components: openapi3.Components{
//			ExtensionProps:  openapi3.ExtensionProps{},
//			Schemas:         nil,
//			Parameters:      nil,
//			Headers:         nil,
//			RequestBodies:   nil,
//			Responses:       nil,
//			SecuritySchemes: nil,
//			Examples:        nil,
//			Links:           nil,
//			Callbacks:       nil,
//		},
//		Info:         nil,
//		Paths:        nil,
//		Security:     nil,
//		Servers:      nil,
//		Tags:         nil,
//		ExternalDocs: nil,
//	}
//	permissions := 0644
//	m, _ := yaml.Marshal(c)
//	err1 := ioutil.WriteFile("file.yaml", m, fs.FileMode(permissions))
//	if err1 != nil {
//		// handle error
//	}
//}
