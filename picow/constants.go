package picow

const (
	DefaultPort = 3000

	CommandGroupConfig = CommandGroup("config")
	CommandGroupInfo   = CommandGroup("info")
	CommandGroupLED    = CommandGroup("led")
	CommandGroupMotion = CommandGroup("motion")

	CommandTypeGet   = CommandType("get")
	CommandTypeSet   = CommandType("set")
	CommandTypeEvent = CommandType("event")
)

type CommandGroup string
type CommandType string
type Command string
