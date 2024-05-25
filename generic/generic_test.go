package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericCache_NewGenericCache(t *testing.T) {
	cache := NewGenericCache[int, string]()

	assert.NotNil(t, cache, "NewGenericCache should return a non-nil cache")
	assert.Empty(t, cache.data, "NewGenericCache should initialize an empty data map")
}

func TestGenericCache_Keys(t *testing.T) {
	cache := NewGenericCache[int, string]()
	cache.Set(1, "one")
	cache.Set(2, "two")
	cache.Set(3, "three")

	keys := cache.Keys()

	assert.Len(t, keys, 3, "Expected 3 keys")

	expectedKeys := []int{1, 2, 3}
	for _, key := range expectedKeys {
		assert.Contains(t, keys, key, "Expected key %d not found", key)
	}
}

func TestGenericCache_Set(t *testing.T) {
	cache := NewGenericCache[int, string]()
	cache.Set(1, "one")

	assert.Len(t, cache.data, 1, "Set should add an entry to the cache")
	assert.Equal(t, "one", cache.data[1], "Set should set the correct value for the given key")
}

func TestGenericCache_Delete(t *testing.T) {
	cache := NewGenericCache[int, string]()
	cache.Set(1, "one")
	cache.Set(2, "two")

	cache.Delete(1)

	assert.Len(t, cache.data, 1, "Delete should remove an entry from the cache")
	assert.NotContains(t, cache.data, 1, "Delete should remove the correct entry")
}

func TestGenericCache_Clear(t *testing.T) {
	cache := NewGenericCache[int, string]()
	cache.Set(1, "one")
	cache.Set(2, "two")

	cache.Clear()

	assert.Empty(t, cache.data, "Clear should remove all entries from the cache")
}

func TestGenericCache_Get(t *testing.T) {
	cache := NewGenericCache[int, string]()
	cache.Set(1, "one")

	value := cache.Get(1)

	assert.Equal(t, "one", value, "Get should return the correct value for the given key")

	nonExistentValue := cache.Get(2)

	assert.Nil(t, nonExistentValue, "Get should return nil for a non-existent key")
}
