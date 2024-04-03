package main

/*
NOTE:
  - Creates a direct connection to a picow device per default
  - A server address can be set to communicate with the main
    server instead with the picow device(s)
  - Wait for motion and (optional) run a command
  - run a command (or run multiple commands at once)
  - always expect json data as a return, as long the request id is not set to "-1"
*/

import (
	"fmt"
	"os"

	"github.com/knackwurstking/picow-led/internal/log"
	"github.com/knackwurstking/picow-led/picow"
)

const (
	// ErrorGeneric - every error not categorized
	ErrorGeneric = 1
	// ErrorArgs - invalid args given (non optional args)
	ErrorArgs = 2
	// ErrorInternal - something went wrong, this is a dev problem :)
	ErrorInternal = 10
	// ErrorUnderConstruction - feature not ready yet
	ErrorUnderConstruction = 100
)

func main() {
	flags := readFlags()
	log.EnableDebug = flags.Debug
	log.Debug("%+v\n", flags)

	// TODO: support for running commands on multiple devices

	req, err := parseArgs(flags.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(ErrorArgs)
	}

	// TODO: create requests per picow address, create a copy of the returned `req` object

	os.Exit(ErrorUnderConstruction)
}

func parseArgs(args []string) (req *picow.Request, err error) {
	// TODO: parse args and return the request object

	return
}
