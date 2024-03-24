package picow

import "fmt"

// TODO: socket communication stuff here

type reader struct {
}

type writer struct {
}

type Server struct {
	Host string
	Port int

	reader reader
	writer writer

	readHandler func()
}

func (s *Server) SetReadHandler(fn func()) error {
	if s.readHandler == nil {
		return fmt.Errorf("a read handler already specified, kill it first")
	}

	s.readHandler = fn

	// TODO: start reading from the server

	return nil
}

func (s *Server) KillReadHandler() {
	s.readHandler = nil

	// TODO: kill read handler
}
