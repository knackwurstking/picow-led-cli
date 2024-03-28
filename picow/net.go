package picow

import (
	"fmt"
	"net"
)

// TODO: socket communication stuff here

type Reader struct {
	conn net.Conn
}

func (r *Reader) Read() (*Response, error) {
	// ...

	return nil, fmt.Errorf("under construction")
}

type Writer struct {
}

type Server struct {
	Host string
	Port int

	reader Reader
	writer Writer

	readHandler func()
}

func (s *Server) GetReader() *Reader {
	return &s.reader
}

func (s *Server) GetWriter() *Writer {
	return &s.writer
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
