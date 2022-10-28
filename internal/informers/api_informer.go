package informers

import (
	"context"
	api_localcache "github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"log"
	"time"
)

func WatchResources(clientSet v1alpha1.APIV1Alpha1Interface, e *APIEventHandler) (cache.Controller, cache.Store) {
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

// APIEventHandler is used to provide functions for resource event handler
type APIEventHandler struct {
	Cache *api_localcache.APILocalCache
}

func (h *APIEventHandler) onAdd(obj interface{}) {
	log.Printf("onAdd called : %v", obj)
}

func (h *APIEventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
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
	log.Printf("onAdd called : %v", obj)
}
