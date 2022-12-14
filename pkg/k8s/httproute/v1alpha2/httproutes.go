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

package v1alpha2

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type HttpRouteInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha2.HTTPRouteList, error)
	Get(ctx context.Context, name string, options metav1.GetOptions) (*v1alpha2.HTTPRoute, error)
	Create(ctx context.Context, route *v1alpha2.HTTPRoute) (*v1alpha2.HTTPRoute, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Update(ctx context.Context, route *v1alpha2.HTTPRoute) (*v1alpha2.HTTPRoute, error)
}

type httpRouteClient struct {
	restClient rest.Interface
	ns         string
}

func (c *httpRouteClient) List(ctx context.Context, opts metav1.ListOptions) (*v1alpha2.HTTPRouteList, error) {
	result := v1alpha2.HTTPRouteList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("HTTPRoutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *httpRouteClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha2.HTTPRoute, error) {
	result := v1alpha2.HTTPRoute{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("HTTPRoutes").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *httpRouteClient) Create(ctx context.Context, httproute *v1alpha2.HTTPRoute) (*v1alpha2.HTTPRoute, error) {
	result := v1alpha2.HTTPRoute{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("HTTPRoutes").
		Body(httproute).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *httpRouteClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("HTTPRoutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (c *httpRouteClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.restClient.Delete().
		Namespace(c.ns).
		Resource("HTTPRoutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *httpRouteClient) Update(ctx context.Context, httproute *v1alpha2.HTTPRoute) (*v1alpha2.HTTPRoute, error) {
	result := v1alpha2.HTTPRoute{}
	err := c.restClient.
		Put().
		Namespace(c.ns).
		Resource("HTTPRoutes").
		Body(httproute).
		Do(ctx).
		Into(&result)

	return &result, err
}
