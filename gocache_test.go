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

	cache.send <- message.Message{
		Action: message.ActionStore,
		Args: map[string]any{
			"key":   "foo",
			"value": "bar",
		},
	}

	r := <-cache.recv

	assert.Equal(t, true, r.Ok)
}
