package fmt

//go:generate ./generate.sh

import (
	"fmt"
	"io"
)

// Printfln ...
func Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Println(fmt.Sprintf(format, a...))
}

// Fprintfln ...
func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, fmt.Sprintf(format, a...))
}
