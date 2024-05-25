package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	gc := NewGoCache()
	gc.data["test"] = "data"
	assert.Equal(t, []string{"test"}, gc.Keys())
	assert.Nil(t, gc.Get("nonexistent"))
}
func TestGet(t *testing.T) {
	gc := NewGoCache()
	gc.data["test"] = "data"
	assert.Equal(t, "data", gc.Get("test"))
	assert.Nil(t, gc.Get("nonexistent"))
}

func TestGetInt(t *testing.T) {
	gc := NewGoCache()
	gc.data["test"] = 42
	val, ok := gc.GetInt("test")
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	_, ok = gc.GetInt("nonexistent")
	assert.False(t, ok)
}

func TestGetInt64(t *testing.T) {
	gc := NewGoCache()
	gc.data["test"] = int64(1234567890)
	val, ok := gc.GetInt64("test")
	assert.True(t, ok)
	assert.Equal(t, int64(1234567890), val)

	_, ok = gc.GetInt64("nonexistent")
	assert.False(t, ok)
}

func TestGetBool(t *testing.T) {
	gc := NewGoCache()
	gc.data["test"] = true
	val, ok := gc.GetBool("test")
	assert.True(t, ok)
	assert.Equal(t, true, val)

	_, ok = gc.GetBool("nonexistent")
	assert.False(t, ok)
}

func TestSet(t *testing.T) {
	gc := NewGoCache()
	gc.Set("test", "data")
	assert.Equal(t, "data", gc.Get("test"))
}

func TestDelete(t *testing.T) {
	gc := NewGoCache()
	gc.Set("test", "data")
	gc.Delete("test")
	assert.Nil(t, gc.Get("test"))
}

func TestClear(t *testing.T) {
	gc := NewGoCache()
	gc.Set("test1", "data1")
	gc.Set("test2", "data2")
	gc.Clear()
	assert.Nil(t, gc.Get("test1"))
	assert.Nil(t, gc.Get("test2"))
}
func TestGetString(t *testing.T) {
	gc := NewGoCache()
	gc.data["test"] = "data"
	val, ok := gc.GetString("test")
	assert.True(t, ok)
	assert.Equal(t, "data", val)
	_, ok = gc.GetString("nonexistent")
	assert.False(t, ok)
}
