package gocache

import (
	"encoding/gob"
	"os"

	"gocache/message"
)

type Server interface {
	HandleMessages()
	SaveToDisk() error
	LoadFromDisk() error
}

type CacheServer struct {
	send            chan message.Response
	recv            chan message.Message
	Data            map[string]any
	PersistanceFile string
	PersistOnAction bool
}

func (server CacheServer) HandleMessages() {
	msg := <-server.recv

	switch msg.Action {
	case message.ActionStatus:
		server.send <- message.Response{
			Ok:    true,
			Value: "OK",
		}
	case message.ActionStore:
		server.Data[msg.Args["key"].(string)] = msg.Args["value"]

		server.send <- message.Response{
			Ok:    true,
			Value: nil,
		}
	case message.ActionGet:
		value, ok := server.Data[msg.Args["key"].(string)]
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

		delete(server.Data, key)

		server.send <- message.Response{
			Ok:    true,
			Value: nil,
		}
	case message.ActionList:
		keys := make([]string, 0)
		for key := range server.Data {
			keys = append(keys, key)
		}
		server.send <- message.Response{
			Ok:    true,
			Value: keys,
		}
	}
}

func (s CacheServer) SaveToDisk() error {
	file, err := os.OpenFile(s.PersistanceFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return gob.NewEncoder(file).Encode(s)
}

func (s CacheServer) LoadFromDisk() error {
	file, err := os.OpenFile(s.PersistanceFile, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	var loadedServer CacheServer

	err = gob.NewDecoder(file).Decode(&loadedServer)

	s.Data = loadedServer.Data
	return err
}

func runCacheServer(server Server) {
	for {
		server.HandleMessages()
	}
}

func StartCache(config GoCacheConfig) CacheClient {
	send := make(chan message.Message, 8)
	recv := make(chan message.Response, 8)

	server := CacheServer{
		send:            recv,
		recv:            send,
		Data:            make(map[string]any),
		PersistanceFile: config.PersistanceFile,
		PersistOnAction: config.PersistOnModification,
	}

	client := CacheClient{
		send,
		recv,
	}

	go runCacheServer(server)

	return client
}
