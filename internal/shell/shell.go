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
	mutex     = sync.Mutex{}
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

func Run(picowDevices []*picow.Net) {
	defer restoreTermState()
	saveTermState()

	for {
		mutex.Lock()
		userCommand := prompt.Input(
			"[picow] ",
			completer,
			prompt.OptionPrefixTextColor(prompt.Blue),
		)
		mutex.Unlock()

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
		for _, device := range picowDevices {
			wg.Add(1)
			go runCommand(&wg, device, cmd, args...)
		}
		wg.Wait()
	}
}

func runCommand(wg *sync.WaitGroup, device *picow.Net, cmd picowcommand.Command, args ...string) {
	defer wg.Done()

	resp, err := cmd.Run(device, args...)
	if err != nil {
		mutex.Lock()
		fmt.Fprintf(os.Stderr, "err: %s %s %s: %s\n", cmd.Group, cmd.Type, cmd.Name, err)
		mutex.Unlock()
		return
	}

	if resp.Error != nil {
		mutex.Lock()
		fmt.Fprintf(os.Stderr, "nerr: response: %s\n", *resp.Error)
		mutex.Unlock()
		return
	}

	// TODO: check server response for data to print
}
