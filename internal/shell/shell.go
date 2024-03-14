package shell

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"

	"github.com/knackwurstking/picow-led/internal/command"
)

var (
	termState *term.State
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

func Run() {
	defer restoreTermState()
	saveTermState()

	for {
		userCommand := prompt.Input(
			"[picow] ",
			completer,
			prompt.OptionPrefixTextColor(prompt.Blue),
		)

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

		cmd := command.New(commandGroup, commandType, commandName)
		if resp, err := cmd.Run(args...); err != nil {
			// TODO: handle error
		} else {
			// TODO: print response back to client (stdout)
		}
	}
}
