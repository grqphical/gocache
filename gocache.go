package gocache

import (
	"errors"

	"gocache/message"
)

var (
	ErrKeyNotFound error = errors.New("gocache: invalid key")
	ErrInvalidType error = errors.New("gocache: invalid type conversion")
)

type CacheServer struct {
	send chan message.Response
	recv chan message.Message
	data map[string]any
}

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

func HandleCacheError(resp message.Response) error {
	if !resp.Ok {
		return resp.Value.(error)
	}

	return nil
}

func Cache(server CacheServer) {
	for {
		msg := <-server.recv

		switch msg.Action {
		case message.ActionStatus:
			server.send <- message.Response{
				Ok:    true,
				Value: "OK",
			}
		case message.ActionStore:
			server.data[msg.Args["key"].(string)] = msg.Args["value"]

			server.send <- message.Response{
				Ok:    true,
				Value: nil,
			}
		case message.ActionGet:
			value, ok := server.data[msg.Args["key"].(string)]
			if !ok {
				server.send <- message.Response{
					Ok:    false,
					Value: ErrKeyNotFound,
				}
			}

			server.send <- message.Response{
				Ok:    true,
				Value: value,
			}
		}
	}
}

func StartCache() CacheClient {
	send := make(chan message.Message, 8)
	recv := make(chan message.Response, 8)

	server := CacheServer{
		send: recv,
		recv: send,
		data: make(map[string]any),
	}

	client := CacheClient{
		send,
		recv,
	}

	go Cache(server)

	return client
}
