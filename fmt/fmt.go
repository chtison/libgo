package fmt

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

// Println ...
func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

// Sprintf ...
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Fprintln ...
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, a...)
}
