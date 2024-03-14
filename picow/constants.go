package picow

const (
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
