package generated

import (
	"fmt"
	"io"
)

// Fmt ...
type Fmt struct {
	Errorf   func(format string, a ...interface{}) error
	Fprint   func(w io.Writer, a ...interface{}) (n int, err error)
	Fprintf  func(w io.Writer, format string, a ...interface{}) (n int, err error)
	Fprintln func(w io.Writer, a ...interface{}) (n int, err error)
	Fscan    func(r io.Reader, a ...interface{}) (n int, err error)
	Fscanf   func(r io.Reader, format string, a ...interface{}) (n int, err error)
	Fscanln  func(r io.Reader, a ...interface{}) (n int, err error)
	Print    func(a ...interface{}) (n int, err error)
	Printf   func(format string, a ...interface{}) (n int, err error)
	Println  func(a ...interface{}) (n int, err error)
	Scan     func(a ...interface{}) (n int, err error)
	Scanf    func(format string, a ...interface{}) (n int, err error)
	Scanln   func(a ...interface{}) (n int, err error)
	Sprint   func(a ...interface{}) string
	Sprintf  func(format string, a ...interface{}) string
	Sprintln func(a ...interface{}) string
	Sscan    func(str string, a ...interface{}) (n int, err error)
	Sscanf   func(str string, format string, a ...interface{}) (n int, err error)
	Sscanln  func(str string, a ...interface{}) (n int, err error)
}

// NewFmt ...
func NewFmt() *Fmt {
	return &Fmt{
		Errorf:   fmt.Errorf,
		Fprint:   fmt.Fprint,
		Fprintf:  fmt.Fprintf,
		Fprintln: fmt.Fprintln,
		Fscan:    fmt.Fscan,
		Fscanf:   fmt.Fscanf,
		Fscanln:  fmt.Fscanln,
		Print:    fmt.Print,
		Printf:   fmt.Printf,
		Println:  fmt.Println,
		Scan:     fmt.Scan,
		Scanf:    fmt.Scanf,
		Scanln:   fmt.Scanln,
		Sprint:   fmt.Sprint,
		Sprintf:  fmt.Sprintf,
		Sprintln: fmt.Sprintln,
		Sscan:    fmt.Sscan,
		Sscanf:   fmt.Sscanf,
		Sscanln:  fmt.Sscanln,
	}
}
