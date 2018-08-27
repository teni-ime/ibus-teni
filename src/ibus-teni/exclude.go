package main

import "sync"

type ExcludeMap struct {
	sync.RWMutex
	m map[string]bool
}

func (e *ExcludeMap) Contains(s string) bool {
	e.RLock()
	defer e.RUnlock()
	if len(e.m) == 0 {
		return false
	}
	_, exist := e.m[s]

	return exist
}

func (e *ExcludeMap) Update(ss []string) {
	e.Lock()
	defer e.Unlock()

	e.m = map[string]bool{}
	for _, s := range ss {
		e.m[s] = true
	}
}
