package gocache

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaves(t *testing.T) {
	s := CacheServer{
		Data:            map[string]any{"foo": "bar"},
		PersistanceFile: "data.bin",
	}

	assert.NoError(t, s.SaveToDisk())

	s1 := CacheServer{
		PersistanceFile: "data.bin",
	}

	assert.NoError(t, s1.LoadFromDisk())

	assert.Equal(t, "bar", s.Data["foo"].(string))

	os.Remove("data.bin")
}
