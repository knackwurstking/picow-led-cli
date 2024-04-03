package picow

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	GroupConfig = Group("config")
	GroupInfo   = Group("info")
	GroupLED    = Group("led")
	GroupMotion = Group("motion")

	TypeSet   = Type("set")
	TypeGet   = Type("get")
	TypeEvent = Type("event")

	IDNoResponse  = ID(-1)
	IDMotionEvent = ID(-2)

	DefaultPort    = 3000
	DefaultEndByte = byte('\n')
)

var (
	Groups = []Group{
		GroupConfig,
		GroupInfo,
		GroupLED,
		GroupMotion,
	}

	Types = []Type{
		TypeSet,
		TypeGet,
		TypeEvent,
	}
)

// Group of command
type Group string

// Type of command
type Type string

// ID of command
type ID int

// Request object for the picow device
type Request struct {
	ID      int      `json:"id"`
	Group   Group    `json:"group"`
	Type    Type     `json:"type"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

// Response object the picow device will respond with
type Response struct {
	ID    int    `json:"id"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

// Server will handle all communication to a picow device
type Server struct {
	host string
	port int

	addr string
	conn net.Conn
}

// NewServer will create a new Server object
func NewServer(host string, port int) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

// GetHost of the current picow device
func (s *Server) GetHost() string {
	return s.host
}

// GetPort of the current picow device
func (s *Server) GetPort() int {
	return s.port
}

// Connect to picow device socket, uses "tcp"
func (s *Server) Connect() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	s.addr = addr
	s.conn = c

	return nil
}

// GetResponse from the picow device
func (s *Server) GetResponse() (*Response, error) {
	// check connection to the picow device
	if s.addr == "" {
		return nil, fmt.Errorf("not connected to server, run connect method first")
	}

	// read data from client
	data := make([]byte, 0)
	chunk := make([]byte, 1)
	for {
		// read byte for byte and check for error
		n, err := s.conn.Read(chunk)
		if err != nil {
			return nil, err
		}

		// break on empty data
		if n == 0 {
			break
		}

		// checking for endbyte
		if chunk[0] == DefaultEndByte {
			break
		}

		// append chunk to data
		data = append(data, chunk...)
	}

	// check data
	if len(data) == 0 {
		return nil, fmt.Errorf("no data")
	}

	resp := Response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &Response{}, nil
}

// Send a request to the picow
func (s *Server) Send(req Request) error {
	// check connection to picow device
	if s.addr == "" {
		return fmt.Errorf("not connected to server, run connect method first")
	}

	// convert request to data
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// write data to client
	n, err := s.conn.Write(append(data, DefaultEndByte))
	if err != nil {
		return err
	} else if n == 0 {
		return fmt.Errorf("no data written to \"%s\"", s.addr)
	}

	return nil
}
