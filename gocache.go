package gocache

import (
	"fmt"

	"gocache/message"
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

func RunCache(server CacheServer) {
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

			fmt.Printf("%+v\n", server.data)

			server.send <- message.Response{
				Ok:    true,
				Value: nil,
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

	go RunCache(server)

	return client
}
