package generated

import (
	"fmt"
	"io"
)

type Fmt struct{}

func NewFmt() *Fmt { return &Fmt{} }

func (_ *Fmt) Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func (_ *Fmt) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, a...)
}

func (_ *Fmt) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

func (_ *Fmt) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, a...)
}

func (_ *Fmt) Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	return fmt.Fscan(r, a...)
}

func (_ *Fmt) Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) {
	return fmt.Fscanf(r, format, a...)
}

func (_ *Fmt) Fscanln(r io.Reader, a ...interface{}) (n int, err error) {
	return fmt.Fscanln(r, a...)
}

func (_ *Fmt) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

func (_ *Fmt) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func (_ *Fmt) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (_ *Fmt) Scan(a ...interface{}) (n int, err error) {
	return fmt.Scan(a...)
}

func (_ *Fmt) Scanf(format string, a ...interface{}) (n int, err error) {
	return fmt.Scanf(format, a...)
}

func (_ *Fmt) Scanln(a ...interface{}) (n int, err error) {
	return fmt.Scanln(a...)
}

func (_ *Fmt) Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

func (_ *Fmt) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func (_ *Fmt) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func (_ *Fmt) Sscan(str string, a ...interface{}) (n int, err error) {
	return fmt.Sscan(str, a...)
}

func (_ *Fmt) Sscanf(str string, format string, a ...interface{}) (n int, err error) {
	return fmt.Sscanf(str, format, a...)
}

func (_ *Fmt) Sscanln(str string, a ...interface{}) (n int, err error) {
	return fmt.Sscanln(str, a...)
}
