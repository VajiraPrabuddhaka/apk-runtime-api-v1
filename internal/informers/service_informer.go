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
	"fmt"
	service_localcache "github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/service"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"time"
)

type ServiceController struct {
	informerFactory informers.SharedInformerFactory
	serviceInformer coreinformers.ServiceInformer
}

// Run starts shared informers and waits for the shared informer cache to
// synchronize.
func (c *ServiceController) Run(stopCh chan struct{}) error {
	// Starts all the shared informers that have been created by the factory so
	// far.
	c.informerFactory.Start(stopCh)
	// wait for the initial synchronization of the local cache.
	if !cache.WaitForCacheSync(stopCh, c.serviceInformer.Informer().HasSynced) {
		return fmt.Errorf("Failed to sync")
	}
	return nil
}

func WatchServices(clientset *kubernetes.Clientset, e *ServiceEventHandler) *ServiceController {
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Hour*24)
	serviceInformer := informerFactory.Core().V1().Services()

	c := &ServiceController{
		informerFactory: informerFactory,
		serviceInformer: serviceInformer,
	}
	serviceInformer.Informer().AddEventHandler(
		// Your custom resource event handlers.
		cache.ResourceEventHandlerFuncs{
			// Called on creation
			AddFunc: e.onAdd,
			// Called on resource update and every resyncPeriod on existing resources.
			UpdateFunc: e.OnUpdate,
			// Called on resource deletion.
			DeleteFunc: e.OnDelete,
		},
	)
	return c
}

type ServiceEventHandler struct {
	Cache *service_localcache.ServiceLocalCache
}

func (h *ServiceEventHandler) onAdd(obj interface{}) {
	v1, _ := obj.(*v1.Service)
	svc := service_localcache.Service{
		Id:          string(v1.UID),
		Namespace:   v1.Namespace,
		Name:        v1.Name,
		Type:        v1.TypeMeta.Kind,
		PortMapping: *service.GeneratePortMapping(v1),
	}

	h.Cache.Update(svc, time.Now().Unix())
}

func (h *ServiceEventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	//Todo perform comparison of oldObj & newObj
	v1, _ := newObj.(*v1.Service)
	svc := service_localcache.Service{
		Id:          string(v1.UID),
		Namespace:   v1.Namespace,
		Name:        v1.Name,
		Type:        v1.TypeMeta.Kind,
		PortMapping: *service.GeneratePortMapping(v1),
	}
	h.Cache.Update(svc, time.Now().Unix())
}

func (h *ServiceEventHandler) OnDelete(obj interface{}) {
	v1, _ := obj.(*v1.Service)
	h.Cache.Delete(string(v1.UID))
}
