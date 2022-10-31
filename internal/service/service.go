package service

import (
	"context"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"k8s.io/client-go/kubernetes"
	"log"

	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetService(namespace string, sid string, clientSet *kubernetes.Clientset) gen.Service {
	servicesClient := clientSet.CoreV1().Services(namespace)

	s, err := servicesClient.Get(context.TODO(), sid, metav1.GetOptions{})

	if err != nil {
		// todo: handle error properly
		return gen.Service{}
	}

	return gen.Service{
		Name:        s.Name,
		Namespace:   s.Namespace,
		Portmapping: GeneratePortMapping(s),
		Type:        s.TypeMeta.Kind,
	}
}

func GetServiceUsage(namespace string, sid string, clientSet *kubernetes.Clientset) ServiceUsageResponse {
	//Todo add logic to find service usage
	return ServiceUsageResponse{
		Count: 0,
		List:  nil,
		Pagination: cache.Pagination{
			Offset:   0,
			Limit:    0,
			Total:    0,
			Next:     "",
			Previous: "",
		},
	}
}

func GetServices(namespace string, offset int, limit int, clientSet *kubernetes.Clientset) ListServicesResponse {
	servicesClient := clientSet.CoreV1().Services(namespace)

	svs, err := servicesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing services : %v", err)
	}
	s, t := processServiceList(svs)
	return ListServicesResponse{
		List:  s,
		Count: len(s),
		Pagination: cache.Pagination{
			Offset:   offset,
			Limit:    limit,
			Total:    t,
			Next:     "todo",
			Previous: "todo",
		},
	}
}

func processServiceList(servicesList *v1.ServiceList) (svs []*gen.Service, total int) {
	var services []*gen.Service
	for _, element := range servicesList.Items {
		var portMapping []gen.PortMapping
		for _, p := range element.Spec.Ports {
			portMapping = append(portMapping, gen.PortMapping{
				Name:       p.Name,
				Protocol:   p.AppProtocol,
				Targetport: p.TargetPort.IntVal,
				Port:       p.Port,
			})
		}
		service := &gen.Service{
			Name:        element.GetName(),
			Namespace:   element.GetNamespace(),
			Type:        element.TypeMeta.Kind,
			Portmapping: &portMapping,
		}
		services = append(services, service)
	}
	//TODO: add proper logic to get total count (handle pagination)
	return services, 6
}

func GeneratePortMapping(svc *v1.Service) *[]gen.PortMapping {
	var portMapping []gen.PortMapping
	for _, p := range svc.Spec.Ports {
		portMapping = append(portMapping, gen.PortMapping{
			Name:       p.Name,
			Protocol:   p.AppProtocol,
			Targetport: p.TargetPort.IntVal,
			Port:       p.Port,
		})
	}
	return &portMapping
}
