package main

import (
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/informers"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/runtime"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/client"
	"net/http"
	"time"
)

func main() {

	s := &runtime.Server{
		ClientSetK8s:      client.GetOutClusterClientSetKubernetes(),
		ClientSetV1alpha1: client.GetOutClusterHttpRouteClientSetV1alpha1(),
		ApiCache:          cache.NewLocalCache(2 * time.Minute),
	}
	h := gen.Handler(s)

	evt := &informers.APIEventHandler{
		Cache: cache.NewLocalCache(60 * time.Minute),
	}

	ctrl, _ := informers.WatchResources(client.GetOutClusterAPIClientSetV1alpha1(), evt)

	go func() {
		stop := make(chan struct{})
		defer close(stop)
		ctrl.Run(stop)
	}()
	http.ListenAndServe(":3000", h)
}
