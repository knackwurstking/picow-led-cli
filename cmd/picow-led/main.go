package main

import (
	"encoding/json"
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
	// ErrorServerError - something went wrong on the server side
	ErrorServerError = 15
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
	log.Debugf("flags=%+v\n", flags)

	for _, subArgs := range subsArgs {
		// parse args for sub
		switch SubCMD(subArgs[0]) {
		case SubCMDRun:
			subFlags, err := flags.ReadSubCMDRun(subArgs[1:])
			if err != nil {
				log.Fatalf(ErrorArgs, "Parse \"%s\" args failed: %s", subArgs[0], err.Error())
			}
			RunCommand(flags.Addr, subFlags, getRequestFromArgs(subFlags.Args))
		case SubCMDOn:
			subFlags, err := flags.ReadSubCMDOn(subArgs[1:])
			if err != nil {
				log.Fatalf(ErrorArgs, "Parse \"%s\" args failed: %s", subArgs[0], err.Error())
			}
			OnEvent(flags.Addr, subFlags)
		default:
			log.Fatalf(ErrorArgs, "Ooops, subcommand \"%s\" not found!", subArgs[0])
		}
	}

	os.Exit(ErrorUnderConstruction)
}

func RunCommand(addr Addr, flags *FlagsSubCMDRun, request *picow.Request) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	defer wg.Done()

	request.ID = flags.ID
	for _, pico := range addr {
		log.Debugf("run command for \"%s\"", pico)

		wg.Add(1)
		go handleRequest(pico, request, flags.PrettyPrint, &wg)
	}

	return &wg
}

func OnEvent(addr Addr, flags *FlagsSubCMDOn) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	defer wg.Done()

	if flags.StartMotion {
		request := &picow.Request{
			ID:      int(picow.IDNoResponse),
			Group:   picow.GroupMotion,
			Type:    picow.TypeEvent,
			Command: "start",
			Args:    make([]string, 0),
		}

		// TODO: send request to server and continue
	}

	// TODO: check if -start-motion flags is set, run command `motion event start` first if true
	// ...

	// TODO: run command: start event, check response for error
	// TODO: and wait for event before return
	// ...

	os.Exit(ErrorUnderConstruction)
	return &wg
}

func handleRequest(pico string, request *picow.Request, prettyResponse bool, wg *sync.WaitGroup) {
	defer wg.Done()

	server := picow.NewServer(pico)
	if err := server.Connect(); err != nil {
		log.Errorf("Connecting to \"%s\" failed: %s", server.GetAddr(), err.Error())
		return
	}

	err := server.Send(request)
	if err != nil {
		log.Errorf("Send request to \"%s\" failed: %s", server.GetAddr(), err.Error())
		return
	}

	if request.ID == int(picow.IDNoResponse) {
		return
	}

	resp, err := server.GetResponse()
	if err != nil {
		log.Errorf("Get response from \"%s\" failed: %s", server.GetAddr(), err.Error())
		return
	}

	if resp.Error != "" {
		if resp.ID != 0 {
			log.Errorf("ID %d: %s: %s", resp.ID, server.GetAddr(), resp.Error)
		} else {
			log.Errorf("%s: %s", server.GetAddr(), resp.ID, resp.Error)
		}
		return
	}

	if resp.Data != nil {
		var data []byte
		var err error
		if prettyResponse {
			data, err = json.MarshalIndent(resp.Data, "", "    ")
		} else {
			data, err = json.Marshal(resp.Data)
		}
		if err != nil {
			log.Debugf("resp.Data=%+v", resp.Data)
			log.Fatalf(ErrorServerError, "Invalid json data from server \"%s\"", server.GetAddr())
		}

		log.Log("%s", string(data))
	}
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
