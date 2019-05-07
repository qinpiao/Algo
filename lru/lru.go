package lru

import (
	"errors"
	"fmt"
	"sync"
)


type LruCache struct {
	cap int
	len int
	l []interface{}
	cache map[interface{}]interface{}
	lock sync.RWMutex
}


func NewLruCache(cap int) *LruCache {
	lru := &LruCache{}
	lru.cap = cap
	lru.len = 0
	lru.l = []interface{}{}
	lru.cache = make(map[interface{}]interface{})
	return lru
}

func (lru *LruCache) Put(key, value interface{}) error {
	if _, OK := lru.cache[key]; OK {
		return errors.New(fmt.Sprintf("Cache already exists: %v", key))
	}
	if lru.len >= lru.cap {
		lru.removeLatest()
	}
	lru.cache[key] = value
	lru.l = append(lru.l, key)
	lru.len ++
	return nil
}

func (lru *LruCache) Get(key interface{}) (interface{}, error) {
	if value, OK := lru.cache[key]; OK {
		lru.refresh(key)
		return value, nil
	}
	return nil, errors.New(fmt.Sprintf("KeyError: %v", key))
}

func (lru *LruCache) Update(key, value interface{}) error{
	if _, OK := lru.cache[key]; OK {
		lru.cache[key] = value
		lru.refresh(key)
		return nil
	}
	return errors.New(fmt.Sprintf("KeyError: %v", key))
}

func (lru *LruCache) removeLatest() {
	key := lru.l[0]
	lru.lock.Lock()
	defer lru.lock.Unlock()
	delete(lru.cache, key)
	lru.l = lru.l[1:]
	lru.len--
}

func (lru *LruCache) refresh(key interface{}) {
	var index int
	for i, v := range lru.l {
		if v == key {
			index = i
			break
		}
	}
	lru.lock.RLock()
	defer lru.lock.RUnlock()
	lru.l = append(lru.l[:index], lru.l[index+1:]...)
	lru.l = append(lru.l, key)

}

func (lru *LruCache) String() {
	lru.lock.RLock()
	defer lru.lock.RUnlock()
	for i:= lru.len-1; i >= 0; i-- {
		fmt.Printf("%s: %d\n", lru.l[i], lru.cache[lru.l[i]])
	}

}