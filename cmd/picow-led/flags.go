package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/knackwurstking/picow-led/picow"
)

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
	// TODO: Add `-id` flag to set a custom id
}

func readFlags() *Flags {
	addr := Addr{}
	debug := false

	flag.Var(&addr, "addr", "picow device address (ip[:port] or hostname[:port])")
	flag.BoolVar(&debug, "debug", debug, "enable debug messages")
	flag.Parse()

	return &Flags{
		Addr:  addr,
		Debug: debug,
		Args:  flag.Args(),
	}
}
