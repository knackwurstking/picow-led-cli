package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/knackwurstking/picow-led/picow"
)

const (
	SubCMDRun = SubCMD("run")
	SubCMDOn  = SubCMD("on")
)

// Sub defines subcommands
type SubCMD string

// FlagsRun subcommand flags
type FlagsSubCMDRun struct {
	ID          int      // ID changes the default command id (the motion id is not allowed)
	Args        []string // Args containing all commandline args besides these already parsed
	PrettyPrint bool     // PrettyPrint enables indentation for response data
}

// FlagsOn subcommand flags
type FlagsSubCMDOn struct {
	StartMotion bool   // StartMotion auto start motion sensor if set to true
	Event       string // Event to start, or wait for ("motion")
}

// Addr contains strings "<ip/hostname>:<port>" for the picow devices to connect to
type Addr []string

// String returns a string with all addresses
func (a Addr) String() string {
	return strings.Join(a, ",")
}

// Set adds a new server
func (a *Addr) Set(value string) error {
	matched, _ := regexp.MatchString("^.+:[0-9]+$", value)
	if !matched {
		// no match means we have to add the default port here
		value = fmt.Sprintf("%s:%d", strings.TrimRight(value, ":"), picow.DefaultPort)
	}

	*a = append(*a, value)

	return nil
}

// Flags holds all flag values
type Flags struct {
	Addr  Addr     // Addr containing the picow server addresses
	Debug bool     // Debug enables debugging messages
	Args  []string // Args containing all commandline args besides these already parsed
}

func NewFlags() *Flags {
	return &Flags{
		Args: make([]string, 0),
	}
}

// Read flags from args
func (flags *Flags) Read() *Flags {
	flag.Var(&flags.Addr, "addr", "picow device address (ip[:port] or hostname[:port])")
	flag.BoolVar(&flags.Debug, "debug", flags.Debug, "enable debug messages")

	flag.Parse()
	flags.Args = flag.Args()

	return flags
}

func (flags *Flags) SplitSubs() ([][]string, error) {
	subsArgs := make([][]string, 0)

	for _, arg := range flags.Args {
		if SubCMD(arg) == SubCMDRun {
			subsArgs = append(subsArgs, []string{arg})
		}

		if len(subsArgs) == 0 {
			return subsArgs, fmt.Errorf("no sub command found!")
		}
		subsArgs[len(subsArgs)-1] = append(subsArgs[len(subsArgs)-1], arg)
	}

	return subsArgs, nil
}

func (*Flags) ReadSubCMDRun(args []string) (*FlagsSubCMDRun, error) {
	cmd := flag.NewFlagSet("run", flag.ExitOnError)
	runFlags := &FlagsSubCMDRun{}

	cmd.IntVar(&runFlags.ID, "id", runFlags.ID, "changes the default id in use")
	cmd.BoolVar(&runFlags.PrettyPrint, "pretty-print", runFlags.PrettyPrint, "pretty prints response data")

	err := cmd.Parse(args)
	runFlags.Args = cmd.Args()
	if runFlags.ID == int(picow.IDMotionEvent) && err == nil {
		err = fmt.Errorf("id \"%d\" not allowed!", picow.IDMotionEvent)
	}

	return runFlags, err
}

func (*Flags) ReadSubCMDOn(args []string) (*FlagsSubCMDOn, error) {
	cmd := flag.NewFlagSet("on", flag.ExitOnError)
	onFlags := &FlagsSubCMDOn{}

	cmd.BoolVar(&onFlags.StartMotion, "start-motion", onFlags.StartMotion, "start motion sensor handling on the pico")

	err := cmd.Parse(args)

	// TODO: check args for <event>
	onFlags.Args = cmd.Args()

	return onFlags, err
}
