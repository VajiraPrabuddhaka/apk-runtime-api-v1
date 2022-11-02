/*
 *  Copyright (c) 2022, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
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

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type HttpRouteV1Alpha1Interface interface {
	HttpRoutes(namespace string) HttpRouteInterface
}

type HttpRouteV1Alpha1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*HttpRouteV1Alpha1Client, error) {
	v1alpha2.AddToScheme(scheme.Scheme)
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha2.GroupName, Version: v1alpha2.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &HttpRouteV1Alpha1Client{restClient: client}, nil
}

func (c *HttpRouteV1Alpha1Client) HttpRoutes(namespace string) HttpRouteInterface {
	return &httpRouteClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}
