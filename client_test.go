package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gocache/message"
)

func TestInitCache(t *testing.T) {
	cache := StartCache()

	cache.send <- message.Message{
		Action: message.ActionStatus,
		Args:   nil,
	}

	r := <-cache.recv

	assert.Equal(t, "OK", r.Value.(string))
}

func TestStore(t *testing.T) {
	cache := StartCache()

	err := cache.Store("foo", "bar")

	assert.Equal(t, nil, err)
}

func TestDelete(t *testing.T) {
	cache := StartCache()
	err := cache.Store("foo", "bar")
	assert.Equal(t, nil, err)

	err = cache.Delete("foo")
	assert.Equal(t, nil, err)

	// Test deleting non-existent key
	err = cache.Delete("baz")
	assert.Equal(t, nil, err)
}

func TestGet(t *testing.T) {
	cache := StartCache()
	err := cache.Store("foo", "bar")
	assert.Equal(t, nil, err)

	value, err := cache.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, "bar", value)

	// Test getting non-existent key
	_, err = cache.Get("baz")
	assert.Error(t, err)
}

func TestGetString(t *testing.T) {
	cache := StartCache()
	err := cache.Store("foo", "bar")
	assert.Equal(t, nil, err)

	value, err := cache.GetString("foo")
	assert.NoError(t, err)
	assert.Equal(t, "bar", value)

	// Test getting non-existent key
	_, err = cache.GetString("baz")
	assert.Error(t, err)
}

func TestGetInt(t *testing.T) {
	cache := StartCache()
	err := cache.Store("foo", 42)
	assert.Equal(t, nil, err)

	value, err := cache.GetInt("foo")
	assert.NoError(t, err)
	assert.Equal(t, 42, value)

	// Test getting non-existent key
	_, err = cache.GetInt("baz")
	assert.Error(t, err)
}

func TestGetFloat(t *testing.T) {
	cache := StartCache()
	err := cache.Store("foo", 3.14)
	assert.Equal(t, nil, err)

	value, err := cache.GetFloat("foo")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, value)

	// Test getting non-existent key
	_, err = cache.GetFloat("baz")
	assert.Error(t, err)
}

func TestGetBytes(t *testing.T) {
	cache := StartCache()
	err := cache.Store("foo", []byte{0x01, 0x02, 0x03})
	assert.Equal(t, nil, err)

	value, err := cache.GetBytes("foo")
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x01, 0x02, 0x03}, value)

	// Test getting non-existent key
	_, err = cache.GetBytes("baz")
	assert.Error(t, err)
}
