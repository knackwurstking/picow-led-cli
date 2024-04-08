package main

import (
	"github.com/knackwurstking/picow-led/picow"
)

type ServerCache struct {
	Data []*picow.Server
}

func (sc *ServerCache) Get(addr string) (*picow.Server, error) {
	for _, server := range sc.Data {
		if server.GetAddr() == addr {
			return server, nil
		}
	}

	server := picow.NewServer(addr)
	err := server.Connect()
	return server, err
}
