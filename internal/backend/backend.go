package backend

import "k8s.io/client-go/kubernetes"

func GetBackend(ns string, name string, clientSet *kubernetes.Clientset) (Backend, error) {
	//todo implement logic to get Backend CR from k8s api server

	return Backend{
		Name: "pet-store",
		Spec: Spec{
			CertificateName: "pet-store",
			Http2Enabled:    false,
			Timeout:         0,
			Credentials: Credentials{
				Type:   "Basic",
				Secret: "secret",
			},
			RetryConfig: RetryConfig{
				Count:      0,
				StatusCode: 0,
			},
			CircuitBreakers: CircuitBreakers{
				MaxConnections:     0,
				MaxRequests:        0,
				MaxPendingRequests: 0,
				MaxRetries:         0,
				MaxConnectionPools: 0,
			},
		},
	}, nil
}
