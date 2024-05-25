package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	gc := NewGoCache()

	gc.data["test"] = "data"

	assert.Equal(t, []string{"test"}, gc.Keys())
}
