package threadsafemap

import (
	"sync"
)

// SafeMap is a thread-safe map with string keys and interface{} values
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

// NewSafeMap creates a new SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

// Get retrieves a value from the map
func (sm *SafeMap) Get(key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

// Set stores a value in the map
func (sm *SafeMap) Set(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// Delete removes a value from the map
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}
