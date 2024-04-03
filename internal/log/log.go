package log

import (
	"fmt"
	"os"
)

var (
	EnableDebug bool = false
)

func Debug(format string, a ...any) {
	if !EnableDebug {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}
