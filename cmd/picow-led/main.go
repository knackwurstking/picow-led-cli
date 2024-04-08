package main

import (
	"encoding/json"
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
	// ErrorServerError - something went wrong on the server side
	ErrorServerError = 15
	// ErrorUnderConstruction - feature not ready yet
	ErrorUnderConstruction = 100
)

var (
	serverCache = &ServerCache{}
)

func main() {
	flags := NewFlags()
	subsArgs, err := flags.Read().SplitSubs()
	if err != nil {
		log.Fatalf(ErrorArgs, "Pasrsing flags failed: %s", err)
	}

	log.EnableDebug = flags.Debug
	log.Debugf("flags=%+v\n", flags)

	// TODO: check flags for `-loop`
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
	for _, a := range addr {
		log.Debugf("run command for \"%s\"", a)

		wg.Add(1)
		func(a string, wg *sync.WaitGroup) {
			defer wg.Done()

			server, err := serverCache.Get(a)
			if err != nil {
				log.Errorf("Server connection for \"%s\" failed: %s", server.GetAddr(), err.Error())
				return
			}

			if err := handleRequest(server, request, flags.PrettyPrint); err != nil {
				log.Errorf("handle request to \"%s\" failed: %s", a, err.Error())
			}
		}(a, &wg)
	}

	return &wg
}

func OnEvent(addr Addr, flags *FlagsSubCMDOn) {
	wg := sync.WaitGroup{}

	if flags.StartMotion {
		// FIXME: Only start once? (in case of -loop is set)
		request := &picow.Request{
			ID:      int(picow.IDNoResponse),
			Group:   picow.GroupMotion,
			Type:    picow.TypeEvent,
			Command: "start",
			Args:    make([]string, 0),
		}

		for _, a := range addr {
			wg.Add(1)
			go func(a string, wg *sync.WaitGroup) {
				defer wg.Done()

				server, err := serverCache.Get(a)
				if err != nil {
					log.Errorf("Server connection for \"%s\" failed: %s", server.GetAddr(), err.Error())
					return
				}

				if err := handleRequest(server, request, false); err != nil {
					log.Errorf("handle request to \"%s\" failed: %s", a, err.Error())
				}
			}(a, &wg)
		}
	}

	// TODO: wait for a motion event, than return (no goroutine, blocking)
	// ...

	os.Exit(ErrorUnderConstruction)
	wg.Wait()
}

func handleRequest(server *picow.Server, request *picow.Request, prettyResponse bool) error {
	// TODO: this could be a problem, i'm not closing existing connections so there is no need to reconnect every time
	//server := picow.NewServer(addr)
	//if err := server.Connect(); err != nil {
	//	return fmt.Errorf("connection failed: %s", err.Error())
	//}

	err := server.Send(request)
	if err != nil {
		return fmt.Errorf("request failed: %s", err.Error())
	}

	if request.ID == int(picow.IDNoResponse) {
		return nil
	}

	resp, err := server.GetResponse()
	if err != nil {
		return fmt.Errorf("get response from \"%s\" failed: %s", server.GetAddr(), err.Error())
	}

	if resp.Error != "" {
		if resp.ID != 0 {
			err = fmt.Errorf("id %d: %s: %s", resp.ID, server.GetAddr(), resp.Error)
		} else {
			err = fmt.Errorf("%s: %s", server.GetAddr(), resp.Error)
		}
		return err
	}

	if resp.Data != nil {
		var data []byte
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

	return nil
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
