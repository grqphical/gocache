package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gocache/message"
)

func TestCache(t *testing.T) {
	cache := StartCache()

	cache.send <- &message.StatusMessage{}

	r := <-cache.recv

	assert.Equal(t, "OK", r.Value.(string))
}
