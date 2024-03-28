package picow

const (
	GroupConfig = Group("config")
	GroupInfo   = Group("info")
	GroupLED    = Group("led")
	GroupMotion = Group("motion")

	TypeSet   = Type("set")
	TypeGet   = Type("get")
	TypeEvent = Type("event")

	IDNoResponse  = ID(-1)
	IDMotionEvent = ID(-2)
)

type Group string
type Type string
type ID int

// TODO:
//  - Request package
//  - Response package
//  - Server struct for communication with a picow device (does not handle multiple devices)
