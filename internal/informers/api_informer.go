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

package informers

import (
	"context"
	api_localcache "github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"time"
)

func WatchAPIs(clientSet v1alpha1.APIV1Alpha1Interface, e *APIEventHandler) (cache.Controller, cache.Store) {
	//h := APIEventHandler{}
	apiStore, apiController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.APIs("default").List(context.TODO(), lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.APIs("default").Watch(context.TODO(), lo)
			},
		},
		&v1alpha1.API{},
		60*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    e.onAdd,
			UpdateFunc: e.OnUpdate,
			DeleteFunc: e.OnDelete,
		},
	)

	return apiController, apiStore
}

// APIEventHandler is used to provide functions for API resource event handler
type APIEventHandler struct {
	Cache *api_localcache.APILocalCache
}

func (h *APIEventHandler) onAdd(obj interface{}) {
	v1, _ := obj.(*v1alpha1.API)

	api := api_localcache.API{
		Id:          string(v1.UID),
		Namespace:   v1.Namespace,
		Name:        v1.Spec.APIDisplayName,
		Context:     v1.Spec.Context,
		Version:     v1.Spec.APIVersion,
		Type:        v1.Spec.APIType,
		CreatedTime: v1.CreationTimestamp.String(),
		UpdatedTime: time.Now().String(),
	}
	h.Cache.Update(api, time.Now().Unix())
}

func (h *APIEventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	//Todo perform comparison of oldObj & newObj
	v1, _ := newObj.(*v1alpha1.API)

	api := api_localcache.API{
		Id:          string(v1.UID),
		Namespace:   v1.Namespace,
		Name:        v1.Spec.APIDisplayName,
		Context:     v1.Spec.Context,
		Version:     v1.Spec.APIVersion,
		Type:        v1.Spec.APIType,
		CreatedTime: v1.CreationTimestamp.String(),
		UpdatedTime: time.Now().String(),
	}
	h.Cache.Update(api, time.Now().Unix())
}

func (h *APIEventHandler) OnDelete(obj interface{}) {
	v1, _ := obj.(*v1alpha1.API)
	h.Cache.Delete(string(v1.UID))
}
