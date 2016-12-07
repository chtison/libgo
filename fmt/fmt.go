package fmt

import (
	"fmt"
	"io"
)

// Printfln formats according to a format specifier and writes to
// standard output with a appended end-of-line (EOL).
func Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Println(fmt.Sprintf(format, a...))
}

// Fprintfln formats according to a format specifier and writes to
// w with a appended end-of-line (EOL).
func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, fmt.Sprintf(format, a...))
}
