package httproute

import (
	"context"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/httproute/gateway/clientset/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	gw_v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func GetHttpRoutes(namespace string, clientSet *v1alpha2.HttpRouteV1Alpha1Client) (routes *gw_v1alpha2.HTTPRouteList, err error) {
	httpRoutes, err := clientSet.HttpRoutes("default").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error listing httproutes : %v", err)
	}
	return httpRoutes, nil
}
