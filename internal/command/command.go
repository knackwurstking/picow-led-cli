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
				Command("led"),
				Command("motion"),
				Command("motion-timeout"),
				Command("pwm-range"),
			},
			TypeGet: {
				Command("led"),
				Command("motion"),
				Command("motion-timeout"),
				Command("pwm-range"),
			},
		},
		GroupInfo: {
			TypeGet: {
				Command("temp"),
				Command("disk-usage"),
				Command("version"),
			},
		},
		GroupLED: {
			TypeSet: {
				Command("duty"),
			},
			TypeGet: {
				Command("duty"),
			},
		},
		GroupMotion: {
			TypeEvent: {
				Command("start"),
				Command("stop"),
			},
		},
	}
)

type Group string
type Type string
type Command string
