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

package httproute

import (
	"context"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/httproute/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	gw_v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func GetHttpRoutes(namespace string, clientSet *v1alpha2.HttpRouteV1Alpha1Client) (routes *gw_v1alpha2.HTTPRouteList, err error) {
	httpRoutes, err := clientSet.HttpRoutes("default").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error listing httproutes : %v", err)
		return nil, err
	}
	return httpRoutes, nil
}

func GetHttpRoute(namespace string, name string, clientSet *v1alpha2.HttpRouteV1Alpha1Client) (*gw_v1alpha2.HTTPRoute, error) {
	httpRoute, err := clientSet.HttpRoutes("default").Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		log.Printf("Error getting httproute : %v", err)
		return nil, err
	}
	return httpRoute, nil
}

func CreateHttpRoute(namespace string, name string, route *gw_v1alpha2.HTTPRoute, clientSet *v1alpha2.HttpRouteV1Alpha1Client) (*gw_v1alpha2.HTTPRoute, error) {
	httpRoute, err := clientSet.HttpRoutes(namespace).Create(context.TODO(), route)

	if err != nil {
		log.Printf("Error creating httproute : %v", err)
		return nil, err
	}
	return httpRoute, err
}

func UpdateHttpRoute(namespace string, name string, route *gw_v1alpha2.HTTPRoute, clientSet *v1alpha2.HttpRouteV1Alpha1Client) (*gw_v1alpha2.HTTPRoute, error) {
	httpRoute, err := clientSet.HttpRoutes(namespace).Update(context.TODO(), route)

	if err != nil {
		log.Printf("Error updating httproute : %v", err)
		return nil, err
	}
	return httpRoute, err
}
