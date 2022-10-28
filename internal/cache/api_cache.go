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

func NewLocalCache(cleanupInterval time.Duration) *APILocalCache {
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
	errUserNotInCache = errors.New("the API isn't in cache")
)

func (lc *APILocalCache) Read(id string) (API, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cu, ok := lc.apis[id]
	if !ok {
		return API{}, errUserNotInCache
	}

	return cu.API, nil
}

func (lc *APILocalCache) Delete(id string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.apis, id)
}
