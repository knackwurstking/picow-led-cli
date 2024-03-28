package main

import "os"

const (
	ErrorGeneric           = 1
	ErrorUnderConstruction = 100
)

func main() {
	// TODO: Handle flags, run a command and print out the response data or error
	// TODO: Support for running commands on multiple devices

	os.Exit(ErrorUnderConstruction)
}
