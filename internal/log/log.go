package log

import (
	"fmt"
	"os"
)

var (
	EnableDebug = false
	PrefixDebug = "debug: "
	PrefixLog   = ""
	PrefixError = "error: "
	PrefixFatal = "error: "
)

// Debug will print a debug message if `EnableDebug` is set to true
func Debugf(format string, a ...any) {
	if !EnableDebug {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, PrefixDebug+format, a...)
}

func Log(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stdout, PrefixLog+format, a...)
}

// Error will just print out an error with prefix
func Errorf(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, PrefixError+format, a...)
}

// Fatal will print out an error message (with prefix) to stderr and exit with code
func Fatalf(code int, format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(code)
}
