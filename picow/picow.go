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

type Group string
type Type string
type ID int

type Request struct {
	ID      int           `json:"id"`
	Group   Group         `json:"group"`
	Type    Type          `json:"type"`
	Command string        `json:"command"`
	Args    []interface{} `json:"args"`
}

type Response struct {
	ID    int    `json:"id"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

type Server struct {
	host string
	port int

	addr string
	conn net.Conn
}

func NewServer(host string, port int) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (s *Server) GetHost() string {
	return s.host
}

func (s *Server) GetPort() int {
	return s.port
}

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

func (s *Server) Read() (*Response, error) {
	if s.addr == "" {
		return nil, fmt.Errorf("not connected to server, run connect method first")
	}

	// TODO: read data from client ()non blocking, until endbyte)

	return nil, fmt.Errorf("under construction")
}

func (s *Server) Write(req Request) error {
	// type checking request.args
	if len(req.Args) > 0 {
		switch req.Args[0].(type) {
		case string:
		case int:
		default:
			return fmt.Errorf("request args list have to be from type string or int")
		}
	}

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
