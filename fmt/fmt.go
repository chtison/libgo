package fmt

import (
	"fmt"
	"io"
	"strings"
)

func Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Println(fmt.Sprintf(format, a...))
}

func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, fmt.Sprintf(format, a...))
}

type Builder struct {
	*strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{
		Builder: &strings.Builder{},
	}
}

func (b *Builder) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(b, a...)
}

func (b *Builder) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(b, a...)
}

func (b *Builder) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(b, format, a...)
}

func (b *Builder) Printfln(format string, a ...interface{}) (int, error) {
	return Fprintfln(b, format, a...)
}
