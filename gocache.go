package gocache

import "sync"

// Cache is a base interface that all GoCache caches implement.
type Cache[K comparable, V any] interface {
	Keys() []K
	Set(K, V)
	Delete(K)
	Clear()
	Get(K) V
}

// GoCache is a basic cache made for most use cases with string keys and dynamic values.
type GoCache struct {
	mu   sync.Mutex
	data map[string]any
}

// NewGoCache creates a new GoCache instance.
func NewGoCache() *GoCache {
	return &GoCache{
		data: map[string]any{},
	}
}

// Keys returns all keys within the store. If there are none, it simply returns an empty array.
//
// Example:
//
//	cache := gocache.NewGoCache()
//	keys := cache.Keys()
func (gc *GoCache) Keys() []string {
	output := make([]string, 0)
	gc.mu.Lock()

	for k := range gc.data {
		output = append(output, k)
	}
	defer gc.mu.Unlock()
	return output
}

// Set sets the value for the given key in the cache.
func (gc *GoCache) Set(key string, value any) {
	gc.mu.Lock()
	gc.data[key] = value
	gc.mu.Unlock()
}

// Delete deletes the value associated with the given key from the cache.
func (gc *GoCache) Delete(key string) {
	gc.mu.Lock()
	delete(gc.data, key)
	gc.mu.Unlock()
}

// Clear clears all the data in the cache.
func (gc *GoCache) Clear() {
	gc.mu.Lock()
	gc.data = make(map[string]any)
	gc.mu.Unlock()
}

// Get retrieves the value associated with the given key from the cache.
// If the key does not exist, it returns nil.
//
//	value := cache.Get("foo")
func (gc *GoCache) Get(key string) any {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	value, ok := gc.data[key]
	if !ok {
		return nil
	}

	return value
}

// GetInt retrieves the value associated with the given key from the cache and returns it as an int.
// If the key does not exist, it returns 0 and false.
//
//	value, exists := cache.GetInt("foo")
//	if !exists {
//		panic("Value does not exist")
//	}
func (gc *GoCache) GetInt(key string) (int, bool) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	value, ok := gc.data[key]
	if !ok {
		return 0, false
	}

	return value.(int), true
}

// GetString retrieves the value associated with the given key from the cache and returns it as a string.
// If the key does not exist, it returns "" and false.
//
//	value, exists := cache.GetString("foo")
//	if !exists {
//		panic("Value does not exist")
//	}
func (gc *GoCache) GetString(key string) (string, bool) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	value, ok := gc.data[key]
	if !ok {
		return "", false
	}

	return value.(string), true
}

// GetInt64 retrieves the value associated with the given key from the cache and returns it as an int64.
// If the key does not exist, it returns 0 and false.
//
//	value, exists := cache.GetInt64("foo")
//	if !exists {
//		panic("Value does not exist")
//	}
func (gc *GoCache) GetInt64(key string) (int64, bool) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	value, ok := gc.data[key]
	if !ok {
		return 0, false
	}

	return value.(int64), true
}

// GetBool retrieves the value associated with the given key from the cache and returns it as a bool.
// If the key does not exist, it returns false and false.
//
//	value, exists := cache.GetBool("foo")
//	if !exists {
//		panic("Value does not exist")
//	}
func (gc *GoCache) GetBool(key string) (bool, bool) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	value, ok := gc.data[key]
	if !ok {
		return false, false
	}

	return value.(bool), true
}
