package main

import "sync"

type CacheMap struct {
	cache map[string]map[string]int
	m     sync.Mutex
}

func (cm *CacheMap) Get(cfg string, record string) int {
	cm.m.Lock()
	defer cm.m.Unlock()
	return cm.cache[cfg][record]
}

func (cm *CacheMap) Set(cfg string, record string, result int) {
	cm.m.Lock()
	defer cm.m.Unlock()
	if cm.cache[cfg] == nil {
		cm.cache[cfg] = map[string]int{
			record: result,
		}
	} else {
		cm.cache[cfg][record] = result
	}
}
