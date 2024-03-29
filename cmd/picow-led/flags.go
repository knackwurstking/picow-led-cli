package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/knackwurstking/picow-led/picow"
)

type Addr []string

func (a Addr) String() string {
	return strings.Join(a, ",")
}

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

type Flags struct {
	Addr flag.Value
}

func readFlags() *Flags {
	var addr flag.Value = &Addr{}

	flag.Var(addr, "addr", "picow device address (ip[:port] or hostname[:port])")
	flag.Parse()

	return &Flags{
		Addr: addr,
	}
}
