package gocache

import (
	"errors"

	"gocache/message"
)

var (
	ErrKeyNotFound          error = errors.New("key not found")
	ErrInvalidType          error = errors.New("invalid type")
	ErrInvalidKeyType       error = errors.New("invalid key type")
	ErrNoPersistanceFileSet error = errors.New("no persistance file has been set")
)

func HandleCacheError(resp message.Response) error {
	if !resp.Ok {
		return resp.Value.(error)
	}

	return nil
}
