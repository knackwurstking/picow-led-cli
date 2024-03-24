package picowcommand

import (
	"fmt"

	"github.com/knackwurstking/picow-led/picow"
)

const (
	GroupConfig = picow.CommandGroupConfig
	GroupInfo   = picow.CommandGroupInfo
	GroupLED    = picow.CommandGroupLED
	GroupMotion = picow.CommandGroupMotion

	TypeSet   = picow.CommandTypeSet
	TypeGet   = picow.CommandTypeGet
	TypeEvent = picow.CommandTypeEvent
)

var (
	Tree = map[picow.CommandGroup]map[picow.CommandType][]picow.Command{
		GroupConfig: {
			TypeSet: {
				picow.Command("led"),
				picow.Command("motion"),
				picow.Command("motion-timeout"),
				picow.Command("pwm-range"),
			},
			TypeGet: {
				picow.Command("led"),
				picow.Command("motion"),
				picow.Command("motion-timeout"),
				picow.Command("pwm-range"),
			},
		},
		GroupInfo: {
			TypeGet: {
				picow.Command("temp"),
				picow.Command("disk-usage"),
				picow.Command("version"),
			},
		},
		GroupLED: {
			TypeSet: {
				picow.Command("duty"),
			},
			TypeGet: {
				picow.Command("duty"),
			},
		},
		GroupMotion: {
			TypeEvent: {
				picow.Command("start"),
				picow.Command("stop"),
			},
		},
	}
)

type Command struct {
	Group picow.CommandGroup
	Type  picow.CommandType
	Name  picow.Command
}

func (c *Command) Run(net *picow.Net, args ...string) (*picow.Response, error) {
	if net != nil {
		return nil, fmt.Errorf("missing server object to connect to")
	}

	// TODO: send data to picow device and await response

	return &picow.Response{}, nil // placeholder return
}

func New(g string, t string, n string) Command {
	return Command{
		Group: picow.CommandGroup(g),
		Type:  picow.CommandType(t),
		Name:  picow.Command(n),
	}
}
