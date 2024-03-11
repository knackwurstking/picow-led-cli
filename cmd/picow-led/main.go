package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
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

func main() {
	defer restoreTermState()
	saveTermState()

	for {
		t := prompt.Input(
			"[picow] ",
			completer,
			prompt.OptionPrefixTextColor(prompt.Blue),
		)

		switch strings.Trim(t, " ") {
		case "exit", "quit", "q":
			exit()
		}

		fmt.Printf("Run command: %s\n", t)
	}
}
