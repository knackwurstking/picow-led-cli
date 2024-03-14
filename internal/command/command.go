package command

const (
	GroupConfig Group = Group("config")
	GroupInfo   Group = Group("info")
	GroupLED    Group = Group("led")
	GroupMotion Group = Group("motion")

	TypeSet   Type = Type("set")
	TypeGet   Type = Type("get")
	TypeEvent Type = Type("event")
)

var (
	Tree map[Group]map[Type][]Name = map[Group]map[Type][]Name{
		GroupConfig: {
			TypeSet: {
				Name("led"),
				Name("motion"),
				Name("motion-timeout"),
				Name("pwm-range"),
			},
			TypeGet: {
				Name("led"),
				Name("motion"),
				Name("motion-timeout"),
				Name("pwm-range"),
			},
		},
		GroupInfo: {
			TypeGet: {
				Name("temp"),
				Name("disk-usage"),
				Name("version"),
			},
		},
		GroupLED: {
			TypeSet: {
				Name("duty"),
			},
			TypeGet: {
				Name("duty"),
			},
		},
		GroupMotion: {
			TypeEvent: {
				Name("start"),
				Name("stop"),
			},
		},
	}
)

type Group string
type Type string
type Name string

type Command struct {
	group Group
	_type Type
	name  Name
}

func (c *Command) Run(args ...string) error {
	// ...

	return nil
}

func New(g string, t string, n string) Command {
	return Command{
		group: Group(g),
		_type: Type(t),
		name:  Name(n),
	}
}
