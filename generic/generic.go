package generic

import "sync"

// GenericCache is a thread-safe cache implementation that stores key-value pairs.
// The keys must be comparable, and the values can be of any type.
type GenericCache[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

// NewGenericCache creates a new instance of GenericCache. It takes in K as the type of keys and V as the type of values to be stored
func NewGenericCache[K comparable, V any]() *GenericCache[K, V] {
	return &GenericCache[K, V]{
		data: map[K]V{},
	}
}

// Keys returns a slice of all the keys in the cache.
func (gc *GenericCache[K, V]) Keys() []K {
	output := make([]K, 0)
	gc.mu.Lock()

	for k := range gc.data {
		output = append(output, k)
	}
	defer gc.mu.Unlock()
	return output
}

// Set sets the value for the given key in the cache.
func (gc *GenericCache[K, V]) Set(key K, value V) {
	gc.mu.Lock()
	gc.data[key] = value
	gc.mu.Unlock()
}

// Delete removes the value associated with the given key from the cache.
func (gc *GenericCache[K, V]) Delete(key K) {
	gc.mu.Lock()
	delete(gc.data, key)
	gc.mu.Unlock()
}

// Clear removes all key-value pairs from the cache.
func (gc *GenericCache[K, V]) Clear() {
	gc.mu.Lock()
	gc.data = make(map[K]V)
	gc.mu.Unlock()
}

// Get retrieves the value associated with the given key from the cache.
// If the key is not found, it returns nil.
func (gc *GenericCache[K, V]) Get(key K) any {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	value, ok := gc.data[key]
	if !ok {
		return nil
	}

	return value
}
