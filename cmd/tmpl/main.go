package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"
	"github.com/chtison/libgo/tmpl"
	"gopkg.in/yaml.v2"
)

var (
	flagYaml = cli.NewFlagStringList('y', "yaml", nil)
)

func init() {
	flagYaml.Usage().Synopsys = "Add data file in yaml format to the template engine's computation"
}

func main() {
	cmd := cli.NewCommand(filepath.Base(os.Args[0]))
	cmd.Usage.Arguments = "[OPTIONS] TEMPLATE [TEMPLATE ...]"
	cmd.Usage.Synopsys = "A CLI for the golang template engine"
	cmd.Function = entrypoint
	cmd.Flags.Add(flagYaml)
	if err := cmd.Execute(os.Args[1:]...); err != nil {
		fmt.Fprintfln(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

func entrypoint(cmd *cli.Command, args ...string) error {
	if len(args) < 1 {
		return cli.Usage(cmd)
	}
	t, err := template.New(args[0]).Funcs(tmpl.Funcs()).ParseFiles(args[0])
	if err != nil {
		return err
	}
	for i := range args[1:] {
		b, err := ioutil.ReadFile(args[i+1])
		if err != nil {
			return err
		}
		fmt.Println(args[i+1])
		if _, err := t.New(args[i+1]).Parse(string(b)); err != nil {
			return err
		}
	}
	m := map[interface{}]interface{}{}
	for _, fileName := range flagYaml.Value {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(data, &m); err != nil {
			return errors.New(fileName + ": " + err.Error())
		}
	}
	if err = t.ExecuteTemplate(os.Stdout, filepath.Base(args[0]), m); err != nil {
		return err
	}
	return nil
}
