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
	"os"

	"github.com/knackwurstking/picow-led/internal/log"
)

const (
	// ErrorGeneric - every error not categorized
	ErrorGeneric = 1
	// ErrorInternal - something went wrong, this is a dev problem :)
	ErrorInternal = 10
	// ErrorUnderConstruction - feature not ready yet
	ErrorUnderConstruction = 100
)

func main() {
	flags := readFlags()
	log.Debug("%+v\n", flags)

	// TODO: support for running commands on multiple devices
	// TODO: init server structs with flags data and run commands
	//       (id, group, type, command, args)

	os.Exit(ErrorUnderConstruction)
}
