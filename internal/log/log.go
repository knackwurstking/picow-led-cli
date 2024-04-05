package log

import (
	"fmt"
	"os"
)

var (
	EnableDebug = false
	PrefixError = "err: "
	PrefixFatal = "err: "
)

// Debug will print a debug message if `EnableDebug` is set to true
func Debugf(format string, a ...any) {
	if !EnableDebug {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, format, a...)
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
