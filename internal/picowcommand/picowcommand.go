package picowcommand

import "github.com/knackwurstking/picow-led/picow"

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
	group picow.CommandGroup
	_type picow.CommandType
	name  picow.Command
}

func (c *Command) Run(args ...string) (*picow.Response, error) {
	// TODO: any - Response type, package: net?
	// ...

	return &picow.Response{}, nil // placeholder return
}

func New(g string, t string, n string) Command {
	return Command{
		group: picow.CommandGroup(g),
		_type: picow.CommandType(t),
		name:  picow.Command(n),
	}
}
