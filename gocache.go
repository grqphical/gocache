package gocache

import (
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
		<-server.recv
		server.send <- message.Response{
			Ok:    true,
			Value: "OK",
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
