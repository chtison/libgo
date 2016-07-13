/*
Package bconv is command line interface for the package baseconverter.

	$ bconv 51966 0123456789 0123456789abcdef
*/
package main

import (
	"fmt"
	"os"

	bc "github.com/chtison/libgo/baseconverter"
)

func main() {
	if len(os.Args) != 4 {
		printUsage()
		return
	}
	number, e1, e2 := bc.BaseToBase(os.Args[1], os.Args[2], os.Args[3])
	if e1 != nil {
		printError(os.Args[2], e1)
		os.Exit(1)
	} else if e2 != nil {
		printError(os.Args[3], e2)
		os.Exit(1)
	}
	fmt.Println(number)
}

func printUsage() {
	fmt.Print(usage)
}

const usage = `usage: bconv {number} {inBase} {toBase}

BaseConverter converts number from base inBase to base toBase.

Example: bconv 51966 0123456789 0123456789abcdef
`

func printError(base string, err error) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(`error: "%s": %s`, base, err))
}
