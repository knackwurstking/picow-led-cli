package picow

// TODO: socket communication stuff here

type Net struct {
	Host string
	Port int
}

func NewNet(host string, port int) *Net {
	return &Net{
		Host: host,
		Port: port,
	}
}
