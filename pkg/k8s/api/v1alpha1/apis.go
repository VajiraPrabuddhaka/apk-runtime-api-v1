/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  WSO2 LLC. licenses this file to you under the Apache License,
 *  Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package v1alpha1

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type APIInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*APIList, error)
	Get(ctx context.Context, name string, options metav1.GetOptions) (*API, error)
	Create(ctx context.Context, api *API) (*API, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Update(ctx context.Context, api *API) (*API, error)
}

type apiClient struct {
	restClient rest.Interface
	ns         string
}

func (c *apiClient) List(ctx context.Context, opts metav1.ListOptions) (*APIList, error) {
	result := APIList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("apis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *apiClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*API, error) {
	result := API{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("apis").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *apiClient) Create(ctx context.Context, api *API) (*API, error) {
	result := API{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("apis").
		Body(api).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *apiClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("apis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (c *apiClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.restClient.Delete().
		Namespace(c.ns).
		Resource("apis").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *apiClient) Update(ctx context.Context, api *API) (*API, error) {
	result := API{}
	err := c.restClient.
		Put().
		Namespace(c.ns).
		Resource("apis").
		Body(api).
		Do(ctx).
		Into(&result)

	return &result, err
}
