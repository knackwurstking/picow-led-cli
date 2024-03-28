package shell

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"

	"github.com/knackwurstking/picow-led/internal/picowcommand"
	"github.com/knackwurstking/picow-led/picow"
)

var (
	termState *term.State
	readMutex = sync.Mutex{}
)

func saveTermState() {
	oldState, err := term.GetState(int(os.Stdin.Fd()))
	if err != nil {
		return
	}
	termState = oldState
}

func restoreTermState() {
	if termState != nil {
		term.Restore(int(os.Stdin.Fd()), termState)
	}
}

func exit() {
	restoreTermState()
	os.Exit(0)
}

func Run(servers ...picow.Server) {
	defer restoreTermState()
	saveTermState()

	readHandlerWG := sync.WaitGroup{}
	defer func() {
		readHandlerWG.Wait()
	}()

	readHandlerWG.Add(1)
	read(&readHandlerWG, servers...)

	for {
		readMutex.Lock()
		userCommand := prompt.Input(
			"[picow] ",
			completer,
			prompt.OptionPrefixTextColor(prompt.Blue),
		)
		readMutex.Unlock()

		switch strings.Trim(userCommand, " ") {
		case "exit", "quit":
			exit()
		}

		cs := make([]string, 0) // command split
		for _, p := range strings.Split(userCommand, " ") {
			if p == "" {
				continue
			}
			cs = append(cs, p)
		}

		var (
			commandGroup string
			commandType  string
			commandName  string
			args         []string
		)

		for i, c := range cs {
			switch i {
			case 0:
				commandGroup = c
			case 1:
				commandType = c
			case 2:
				commandName = c
			default:
				args = cs[i:]
			}
		}

		cmd := picowcommand.New(commandGroup, commandType, commandName)

		wg := sync.WaitGroup{}
		for _, server := range servers {
			wg.Add(1)
			go write(&wg, server, cmd, args...)
		}
		wg.Wait()
	}
}

func read(wg *sync.WaitGroup, servers ...picow.Server) {
	defer wg.Done()

	wg2 := sync.WaitGroup{}
	for _, server := range servers {
		wg2.Add(1)
		go func(wg *sync.WaitGroup, server picow.Server) {
			defer wg.Done()

			reader := server.GetReader()
			response, err := reader.ReadResponse()
			if err != nil {
				readMutex.Lock()
				fmt.Fprintf(os.Stderr, "err: %s\n", err.Error())
				readMutex.Unlock()
				return
			}

			if response.Error != nil {
				if *response.Error != "" {
					readMutex.Lock()
					fmt.Fprintf(os.Stderr, "err: %s\n", *response.Error)
					readMutex.Unlock()
					return
				}
			}

			// TODO: check for data and ID, then print data to stdout
		}(&wg2, server)
	}

	wg.Wait()

	// TODO: read from pico devices and write to stdout/stderr
}

func write(wg *sync.WaitGroup, server picow.Server, cmd picowcommand.Command, args ...string) {
	defer wg.Done()

	err := cmd.Write(server.GetWriter(), args...)
	if err != nil {
		readMutex.Lock()
		fmt.Fprintf(os.Stderr, "err: %s %s %s: %s\n", cmd.Group, cmd.Type, cmd.Name, err)
		readMutex.Unlock()
	}
}
