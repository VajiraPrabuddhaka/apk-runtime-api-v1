/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

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
