package main

import (
	"fmt"
	"os"
)

const (
	ErrorGeneric           = 1
	ErrorInternal          = 10
	ErrorUnderConstruction = 100
)

func main() {
	flags := readFlags()
	fmt.Fprintf(os.Stderr, "%+v", flags)

	// TODO: support for running commands on multiple devices
	// TODO: init server structs with flags data and run commands
	//       (id, group, type, command, args)

	os.Exit(ErrorUnderConstruction)
}
