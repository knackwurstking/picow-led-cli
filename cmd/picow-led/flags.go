package main

import (
	"flag"
	"fmt"
	"os"
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
	fmt.Println(value)
	matched, err := regexp.MatchString("^.+:[0-9]+$", value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: reqular expression failed: %s", err.Error())
		os.Exit(ErrorInternal)
	}

	if !matched {
		value = fmt.Sprintf("%s:%d", strings.TrimRight(value, ":"), picow.DefaultPort)
	}

	*a = append(*a, value)

	return nil
}

// Flags holds all flag values
type Flags struct {
	Addr flag.Value // Addr containing the picow server addresses
	Args []string   // Args containing all commandline args besides these already parsed
}

func readFlags() *Flags {
	var addr flag.Value = &Addr{}

	flag.Var(addr, "addr", "picow device address (ip[:port] or hostname[:port])")
	flag.Parse()

	return &Flags{
		Addr: addr,
		Args: flag.Args(),
	}
}
