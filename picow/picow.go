package picow

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

	DefaultPort = 3000
)

type Group string
type Type string
type ID int

// TODO:
//  - Request package
//  - Response package
//  - Server struct for communication with a picow device (does not handle
//    multiple devices)

type Request[T int | string] struct {
	ID      int    `json:"id"`
	Group   Group  `json:"group"`
	Type    Type   `json:"type"`
	Command string `json:"command"`
	Args    []T    `json:"args"`
}

type Response struct {
	ID    int    `json:"id"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

type Server struct {
	host string
	port int
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
