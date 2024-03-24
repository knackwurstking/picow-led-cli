package main

import (
	"github.com/knackwurstking/picow-led/internal/shell"
	"github.com/knackwurstking/picow-led/picow"
)

var (
	picowDevices []picow.Server = []picow.Server{
		{Host: "127.0.0.1", Port: picow.DefaultPort},
	}
)

func main() {
	shell.Run(picowDevices...)
}
