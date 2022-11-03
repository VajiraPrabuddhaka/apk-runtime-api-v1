/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package cache

import (
	"errors"
	"log"
	"sync"
	"time"
)

type API struct {
	Id          string `json:"id"`
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	Context     string `json:"context"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type ListAPIsResponse struct {
	Count      int         `json:"count"`
	List       []API       `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type CachedAPI struct {
	API
	expireAtTimestamp int64
}

type APILocalCache struct {
	stop chan struct{}

	wg   sync.WaitGroup
	mu   sync.RWMutex
	apis map[string]CachedAPI
}

func NewAPILocalCache(cleanupInterval time.Duration) *APILocalCache {
	lc := &APILocalCache{
		apis: make(map[string]CachedAPI),
		stop: make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)
	return lc
}

func (lc *APILocalCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for uid, cu := range lc.apis {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.apis, uid)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *APILocalCache) stopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *APILocalCache) Update(u API, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.apis[u.Id] = CachedAPI{
		API:               u,
		expireAtTimestamp: expireAtTimestamp,
	}
	log.Printf("Cache updated successfully.. cache: %v", lc.apis)
}

var (
	errAPINotInCache = errors.New("the API isn't in cache")
)

func (lc *APILocalCache) Read(id string) (API, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cu, ok := lc.apis[id]
	if !ok {
		return API{}, errAPINotInCache
	}

	return cu.API, nil
}

func (lc *APILocalCache) ReadAll() ([]API, error) {
	var apis []API

	for _, api := range lc.apis {
		apis = append(apis, api.API)
	}
	return apis, nil
}

func (lc *APILocalCache) APIsSearch(offset *int, limit *int, name *string, namespace *string, sortBy *string,
	sortOrder *string) (*ListAPIsResponse, error) {
	//Todo : error handling
	s, _ := lc.ReadAll()
	apis, count, pagination, _ := GenResultWithPagination(interface{}(s), "API", offset, limit, sortBy, *sortOrder)
	return &ListAPIsResponse{
		Count:      count,
		List:       apis.([]API),
		Pagination: pagination,
	}, nil
}

func (lc *APILocalCache) Delete(id string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.apis, id)
}
