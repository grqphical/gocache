package gocache

import "sync"

type Cache[K any] interface {
	Keys() []K
}

type GoCache struct {
	mu   sync.Mutex
	data map[string]interface{}
}

func NewGoCache() *GoCache {
	return &GoCache{
		data: map[string]interface{}{},
	}
}

func (gc *GoCache) Keys() []string {
	output := make([]string, 0)
	gc.mu.Lock()

	for k := range gc.data {
		output = append(output, k)
	}
	defer gc.mu.Unlock()
	return output
}
