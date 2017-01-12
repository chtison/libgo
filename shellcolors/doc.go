/*
Package shellcolors is a simple interface for terminal Select Graphic
Rendition (SGR).

https://en.wikipedia.org/wiki/ANSI_escape_code#graphics

Package shellcolors lets you modify style of you outputted text in a terminal
like the boldness or the background color.
However, some SGR codes might not be supported by your terminal.

	package main

	import (
		"fmt"
		sc "github.com/chtison/libgo/shellcolors"
	)

	func main() {
		fmt.Printf("%sHello %sWorld %s!%s\n",
			sc.NewWithColor(32, sc.Bold),
			sc.NewWithColor(205, sc.Reset),
			sc.New(sc.BgWhite, sc.Red, sc.Negative),
			sc.New(sc.Reset))
	}


The type ShellColor implements the interface fmt.Stringer.

There are 3 primary functions to handle ShellColor type:

	func New(codes ...CodeSGR) *ShellColor
	func (self *ShellColor) Add(codes ...CodeSGR) *ShellColor
	func (self *ShellColor) String() string

*/
package shellcolors
