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
			if server.IsConnected() {
				return server, nil
			} else {
				err := server.Connect()
				return server, err
			}
		}
	}

	server := picow.NewServer(addr)
	err := server.Connect()
	sc.Data = append(sc.Data, server)
	return server, err
}
