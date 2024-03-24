package main

import (
	"github.com/knackwurstking/picow-led/internal/shell"
	"github.com/knackwurstking/picow-led/picow"
)

var (
	picowDevices []*picow.Net = []*picow.Net{
		picow.NewNet("127.0.0.1", picow.DefaultPort),
	}
)

func main() {
	shell.Run(picowDevices)
}
