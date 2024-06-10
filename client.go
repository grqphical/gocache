package gocache

import (
	"gocache/message"
)

type CacheClient struct {
	send chan message.Message
	recv chan message.Response
}

func (c *CacheClient) Store(key string, value any) error {
	c.send <- message.Message{
		Action: message.ActionStore,
		Args: map[string]any{
			"key":   key,
			"value": value,
		},
	}

	return HandleCacheError(<-c.recv)
}

func (c *CacheClient) Delete(keys ...string) error {
	for _, key := range keys {
		c.send <- message.Message{
			Action: message.ActionDelete,
			Args: map[string]any{
				"key": key,
			},
		}
	}

	resp := <-c.recv
	return HandleCacheError(resp)
}

func (c *CacheClient) Get(key string) (any, error) {
	c.send <- message.Message{
		Action: message.ActionGet,
		Args: map[string]any{
			"key": key,
		},
	}

	resp := <-c.recv
	err := HandleCacheError(resp)
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (c *CacheClient) GetString(key string) (string, error) {
	c.send <- message.Message{
		Action: message.ActionGet,
		Args: map[string]any{
			"key": key,
		},
	}

	resp := <-c.recv
	err := HandleCacheError(resp)
	if err != nil {
		return "", err
	}

	value, ok := resp.Value.(string)
	if !ok {
		return "", ErrInvalidType
	}

	return value, nil
}

func (c *CacheClient) GetInt(key string) (int, error) {
	c.send <- message.Message{
		Action: message.ActionGet,
		Args: map[string]any{
			"key": key,
		},
	}

	resp := <-c.recv
	err := HandleCacheError(resp)
	if err != nil {
		return 0, err
	}

	value, ok := resp.Value.(int)
	if !ok {
		return 0, ErrInvalidType
	}

	return value, nil
}

func (c *CacheClient) GetFloat(key string) (float64, error) {
	c.send <- message.Message{
		Action: message.ActionGet,
		Args: map[string]any{
			"key": key,
		},
	}

	resp := <-c.recv
	err := HandleCacheError(resp)
	if err != nil {
		return 0.0, err
	}

	value, ok := resp.Value.(float64)
	if !ok {
		return 0.0, ErrInvalidType
	}

	return value, nil
}

func (c *CacheClient) GetBytes(key string) ([]byte, error) {
	c.send <- message.Message{
		Action: message.ActionGet,
		Args: map[string]any{
			"key": key,
		},
	}

	resp := <-c.recv
	err := HandleCacheError(resp)
	if err != nil {
		return nil, err
	}

	value, ok := resp.Value.([]byte)
	if !ok {
		return nil, ErrInvalidType
	}

	return value, nil
}
