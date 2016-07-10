# Golang package shellcolors

Package shellcolors is a simple interface for terminal Select Graphic Rendition (SGR).

#### Install this package + the cli frontend of this package: [sc](cmd/sc)
```
$ go get -v github.com/chtison/libgo/shellcolors/...
```

Package shellcolors lets you modify style of you outputed text in a terminal
like the boldness or the background color.
However, some SGR codes might not be supported by your terminal.

https://en.wikipedia.org/wiki/ANSI_escape_code#graphics

```go
package main

import (
	"fmt"
	sc "github.com/chtison/libgo/shellcolors"
)

func main() {
	fmt.Println(fmt.Sprintf("%sHello World !%s",
		sc.New(sc.Bold, sc.Green),
		sc.New(sc.Reset),
	))
}
```


The type ShellColor implements the [fmt.Stringer](https://golang.org/pkg/fmt/#Stringer) interface.

#### There are 3 primary functions to handle ShellColor type:

```go
// Create a new ShellColor "object"
func New(codes ...CodeSGR) *ShellColor { }
// Add a parameter to the ShellColor "object"
func (self *ShellColor) Add(codes ...CodeSGR) *ShellColor { }
// Get back a string formatted for controlling terminal parameters
func (self *ShellColor) String() string { }
```

#### And functions designed for adding colors
```go
// Add font color (256 colors)
func (self *ShellColor) Color(color uint8) *ShellColor { }
// Add background color (256 colors)
func (self *ShellColor) BgColor(color uint8) *ShellColor { }
// Add font color (RGB not always supported)
func (self *ShellColor) ColorRGB(red, green, blue uint8) *ShellColor { }
// Add background color (RGB not always supported)
func (self *ShellColor) BgColorRGB(red, green, blue uint8) *ShellColor { }

func NewWithColor(color uint8, codes ...CodeSGR) *ShellColor { }
func NewWithColorRGB(red, green, blue uint8, codes ...CodeSGR) *ShellColor { }
func NewWithBgColor(color uint8, codes ...CodeSGR) *ShellColor { }
func NewWithBgColorRGB(red, green, blue uint8, codes ...CodeSGR) *ShellColor { }
```
