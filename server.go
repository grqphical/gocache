package gocache

import "gocache/message"

type CacheServer struct {
	send chan message.Response
	recv chan message.Message
	data map[string]any
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
		case message.ActionDelete:
			key := msg.Args["key"].(string)

			delete(server.data, key)

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

	go Cache(server)

	return client
}
