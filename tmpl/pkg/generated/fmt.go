package generated

import (
	"fmt"
	"io"
)

type Fmt struct{}

func NewFmt() *Fmt { return &Fmt{} }

func (*Fmt) Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func (*Fmt) Fprint(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, a...)
}

func (*Fmt) Fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, a...)
}

func (*Fmt) Fprintln(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprintln(w, a...)
}

func (*Fmt) Fscan(r io.Reader, a ...interface{}) (int, error) {
	return fmt.Fscan(r, a...)
}

func (*Fmt) Fscanf(r io.Reader, format string, a ...interface{}) (int, error) {
	return fmt.Fscanf(r, format, a...)
}

func (*Fmt) Fscanln(r io.Reader, a ...interface{}) (int, error) {
	return fmt.Fscanln(r, a...)
}

func (*Fmt) Print(a ...interface{}) (int, error) {
	return fmt.Print(a...)
}

func (*Fmt) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a...)
}

func (*Fmt) Println(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

func (*Fmt) Scan(a ...interface{}) (int, error) {
	return fmt.Scan(a...)
}

func (*Fmt) Scanf(format string, a ...interface{}) (int, error) {
	return fmt.Scanf(format, a...)
}

func (*Fmt) Scanln(a ...interface{}) (int, error) {
	return fmt.Scanln(a...)
}

func (*Fmt) Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

func (*Fmt) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func (*Fmt) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func (*Fmt) Sscan(str string, a ...interface{}) (int, error) {
	return fmt.Sscan(str, a...)
}

func (*Fmt) Sscanf(str string, format string, a ...interface{}) (int, error) {
	return fmt.Sscanf(str, format, a...)
}

func (*Fmt) Sscanln(str string, a ...interface{}) (int, error) {
	return fmt.Sscanln(str, a...)
}
