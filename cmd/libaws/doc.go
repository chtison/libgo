package main

import (
	"github.com/chtison/libgo/cli"
)

var cmdDoc = cli.NewCommand("doc")

func init() {
	cmdDoc.Usage.Synopsys = "Open the webpage reference of AWS CloudFormation resource types"
	cmdDoc.Function = cmdDocFunction
}

func cmdDocFunction(cmd *cli.Command, args ...string) error {
	if len(args) < 1 {
		return cli.Usage(cmd, args...)
	}
	return nil
}
