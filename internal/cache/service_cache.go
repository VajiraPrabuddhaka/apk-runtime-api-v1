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

package cache

import (
	"errors"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"log"
	"sync"
	"time"
)

type Service struct {
	Id          string `json:"id"`
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	PortMapping []gen.PortMapping
}

type ListServicesResponse struct {
	Count      int         `json:"count"`
	List       []Service   `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type CachedService struct {
	Service
	expireAtTimestamp int64
}

type ServiceLocalCache struct {
	stop chan struct{}

	wg       sync.WaitGroup
	mu       sync.RWMutex
	services map[string]CachedService
}

func NewServiceLocalCache(cleanupInterval time.Duration) *ServiceLocalCache {
	lc := &ServiceLocalCache{
		services: make(map[string]CachedService),
		stop:     make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)
	return lc
}

func (lc *ServiceLocalCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for uid, cu := range lc.services {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.services, uid)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *ServiceLocalCache) stopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *ServiceLocalCache) Update(u Service, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.services[u.Id] = CachedService{
		Service:           u,
		expireAtTimestamp: expireAtTimestamp,
	}
	log.Printf("Cache updated successfully.. cache: %v", lc.services)
}

var (
	errServiceNotInCache = errors.New("the Service isn't in cache")
)

func (lc *ServiceLocalCache) Read(id string) (Service, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cu, ok := lc.services[id]
	if !ok {
		return Service{}, errServiceNotInCache
	}

	return cu.Service, nil
}

func (lc *ServiceLocalCache) ReadAll() ([]Service, error) {
	var svs []Service

	for _, svc := range lc.services {
		svs = append(svs, svc.Service)
	}
	return svs, nil
}

func (lc *ServiceLocalCache) ServicesSearch(offset *int, limit *int, name *string, namespace *string, sortBy *string,
	sortOrder *string) (*ListServicesResponse, error) {
	//Todo : error handling
	s, _ := lc.ReadAll()
	svs, count, pagination, _ := GenResultWithPagination(interface{}(s), "Service", offset, limit, sortBy, *sortOrder)
	return &ListServicesResponse{
		Count:      count,
		List:       svs.([]Service),
		Pagination: pagination,
	}, nil
}

func (lc *ServiceLocalCache) Delete(id string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.services, id)
}
