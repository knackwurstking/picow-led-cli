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
	Data map[Group]map[Type][]Command = map[Group]map[Type][]Command{
		GroupConfig: {
			TypeSet: {
				NewCommand(GroupConfig, TypeSet, "led"),
				NewCommand(GroupConfig, TypeSet, "motion"),
				NewCommand(GroupConfig, TypeSet, "motion-timeout"),
				NewCommand(GroupConfig, TypeSet, "pwm-range"),
			},
			TypeGet: {
				NewCommand(GroupConfig, TypeGet, "led"),
				NewCommand(GroupConfig, TypeGet, "motion"),
				NewCommand(GroupConfig, TypeGet, "motion-timeout"),
				NewCommand(GroupConfig, TypeGet, "pwm-range"),
			},
		},
		GroupInfo: {
			TypeGet: {
				NewCommand(GroupInfo, TypeGet, "temp"),
				NewCommand(GroupInfo, TypeGet, "disk-usage"),
				NewCommand(GroupInfo, TypeGet, "version"),
			},
		},
		GroupLED: {
			TypeSet: {
				NewCommand(GroupLED, TypeSet, "duty"),
			},
			TypeGet: {
				NewCommand(GroupLED, TypeGet, "duty"),
			},
		},
		GroupMotion: {
			TypeEvent: {
				NewCommand(GroupMotion, TypeEvent, "start"),
				NewCommand(GroupMotion, TypeEvent, "stop"),
			},
		},
	}
)

type Group string
type Type string

type Command struct {
	Group Group
	Type  Type
	Name  string
}

func NewCommand(group Group, _type Type, command string) Command {
	return Command{
		Group: group,
		Type:  _type,
		Name:  command,
	}
}

func (c Command) String() string {
	return c.Name
}

func (c Command) Run() error {
	// TODO: run command from group of type...

	return nil
}
