package log

import (
	"fmt"
	"os"
)

func Debug(format string, a ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}
