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

	wg := sync.WaitGroup{}
	for _, subArgs := range subsArgs {
		// TODO: parse args for sub

		// TODO: run command in a goroutine, for each address (this should replace the old section after this for loop)
	}

	//// parse args
	//req, err := parseArgs(flags.Args) // TODO: remove this
	//if err != nil {
	//	log.Fatalf(ErrorArgs, "Parsing args failed %s", err)
	//}

	//// send request to picow devices
	//wg := sync.WaitGroup{}
	//for _, addr := range flags.Addr {
	//	wg.Add(1)
	//	go func(addr string, wg *sync.WaitGroup) {
	//		defer wg.Done()
	//		// TODO: send request to server
	//		// ...
	//	}(addr, &wg)
	//}

	//// start read response handler
	//// TODO: read response from server (only if id is not -1)
	//// ...

	wg.Wait()

	// check and output error, or data in json format
	// TODO: print out results

	os.Exit(ErrorUnderConstruction)
}

func parseArgs(args []string) (req *picow.Request, err error) {
	if len(args) < 3 {
		return req, fmt.Errorf("wrong args: <group> <command-type> <command> [<args> ...]")
	}

	group := picow.Group("")
	for _, g := range picow.Groups {
		if g == picow.Group(args[0]) {
			group = g
			break
		}
	}
	if group == "" {
		return req, fmt.Errorf("group not exists: %s", group)
	}

	_type := picow.Type("")
	for _, t := range picow.Types {
		if t == picow.Type(args[1]) {
			_type = t
			break
		}
	}
	if _type == "" {
		return req, fmt.Errorf("command type not exists: %s", _type)
	}

	req = &picow.Request{
		ID:      0,
		Group:   picow.Group(args[0]),
		Type:    picow.Type(args[1]),
		Command: args[2],
		Args:    make([]string, 0),
	}

	return req, err
}
