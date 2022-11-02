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
		ApiCache:          cache.NewAPILocalCache(2 * time.Minute),
		ServiceCache:      cache.NewServiceLocalCache(2 * time.Minute),
	}
	h := gen.Handler(s)

	evtApi := &informers.APIEventHandler{
		Cache: s.ApiCache,
	}

	evtService := &informers.ServiceEventHandler{
		Cache: s.ServiceCache,
	}

	ctrlApi, _ := informers.WatchAPIs(client.GetOutClusterAPIClientSetV1alpha1(), evtApi)

	ctrlSvc := informers.WatchServices(client.GetOutClusterClientSetKubernetes(), evtService)

	go func() {
		stop := make(chan struct{})
		defer close(stop)
		ctrlSvc.Run(stop)
		ctrlApi.Run(stop)
	}()
	http.ListenAndServe(":3000", h)
}
