package main

import (
	"os"
	"path/filepath"

	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"
)

func main() {
	cmd := cli.NewCommand(filepath.Base(os.Args[0]))
	cmd.Usage.Synopsys = "A CLI for manipulating AWS"
	cmd.Function = entrypoint
	cmd.Children.Add(cmdDoc)
	if err := cmd.Execute(os.Args[1:]...); err != nil {
		fmt.Fprintfln(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

func entrypoint(cmd *cli.Command, args ...string) error {
	return cli.Usage(cmd, args...)
}
