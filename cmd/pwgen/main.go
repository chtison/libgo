package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chtison/libgo/pwgen"
)

const usage = `usage: %s LENGTH CHARSET
Pwgen generates a string of LENGTH characters from CHARSET and outputs it to the standard output.

Example:
$ %s 16 abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
x0aPdDriwsXskDCJ
`

func ftUsage() {
	fmt.Printf(usage, os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		ftUsage()
	}
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing LENGTH: %s\n", err)
		os.Exit(2)
	}
	passwd := pwgen.Generate(uint(n), os.Args[2])
	fmt.Println(passwd)
}
