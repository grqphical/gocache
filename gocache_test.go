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

func TestGetString(t *testing.T) {
	cache := StartCache()

	err := cache.Store("foo", "bar")

	assert.Equal(t, nil, err)

	value, err := cache.GetString("foo")
	assert.NoError(t, err)

	assert.Equal(t, "bar", value)

	// Test key that doesn't exist
	_, err = cache.GetString("baz")
	assert.Error(t, err)
}
