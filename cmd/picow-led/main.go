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
	"sync"

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
	flags := NewFlags()
	subsArgs, err := flags.Read().SplitSubs()
	if err != nil {
		log.Fatalf(ErrorArgs, "Pasrsing flags failed: %s", err)
	}

	log.EnableDebug = flags.Debug
	log.Debugf("%+v\n", flags)

	for _, subArgs := range subsArgs {
		// parse args for sub
		switch SubCMD(subArgs[0]) {
		case SubCMDRun:
			subFlags, err := flags.ReadSubCMDRun(subArgs[1:])
			if err != nil {
				log.Fatalf(ErrorArgs, "Parse \"%s\" args failed: %s", subArgs[0], err.Error())
			}
			runCommand(flags.Addr, subFlags, getRequestFromArgs(subFlags.Args))
		case SubCMDOn:
			subFlags, err := flags.ReadSubCMDOn(subArgs[1:])
			if err != nil {
				log.Fatalf(ErrorArgs, "Parse \"%s\" args failed: %s", subArgs[0], err.Error())
			}
			onEvent(flags.Addr, subFlags)
		default:
			log.Fatalf(ErrorArgs, "Ooops, subcommand \"%s\" not found!", subArgs[0])
		}
	}

	os.Exit(ErrorUnderConstruction)
}

func getRequestFromArgs(args []string) (req *picow.Request) {
	if len(args) < 3 {
		log.Fatalf(ErrorArgs, "Wrong ARGS: <group> <command-type> <command> [<args> ...]")
	}

	group := picow.Group("")
	for _, g := range picow.Groups {
		if g == picow.Group(args[0]) {
			group = g
			break
		}
	}
	if group == "" {
		log.Fatalf(ErrorArgs, "Group \"%s\" not exists!", group)
	}

	_type := picow.Type("")
	for _, t := range picow.Types {
		if t == picow.Type(args[1]) {
			_type = t
			break
		}
	}
	if _type == "" {
		log.Fatalf(ErrorArgs, "Command type \"%s\" not exists!", _type)
	}

	req = &picow.Request{
		ID:      0,
		Group:   picow.Group(args[0]),
		Type:    picow.Type(args[1]),
		Command: args[2],
		Args:    make([]string, 0),
	}

	return req
}

func runCommand(addr Addr, subArgs *FlagsSubCMDRun, request *picow.Request) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	defer wg.Done()
	// TODO: run command / send request to server and print out the response

	os.Exit(ErrorUnderConstruction)

	return &wg
}

func onEvent(addr Addr, subArgs *FlagsSubCMDOn) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	defer wg.Done()
	// TODO: get request object from flags

	// TODO: run command: start event, check response for error
	// TODO: and wait for event before return

	os.Exit(ErrorUnderConstruction)

	return &wg
}
