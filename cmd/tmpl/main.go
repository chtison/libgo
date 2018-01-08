package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"
	"gopkg.in/yaml.v2"
)

var (
	flagYaml                 = cli.NewFlagStringList('y', "yaml", nil)
	flagListDefinedTemplates = cli.NewFlagBool('l', "list", false)
)

func init() {
	flagYaml.Usage().Synopsys = "Add data file in yaml format to the template engine's computation"
	flagListDefinedTemplates.Usage().Synopsys = "List defined templates then stop"
}

func main() {
	cmd := cli.NewCommand(filepath.Base(os.Args[0]))
	cmd.Usage.Arguments = "[OPTIONS] TEMPLATE [TEMPLATE ...]"
	cmd.Usage.Synopsys = "A CLI for the golang template engine"
	cmd.Function = entrypoint
	cmd.Flags.Add(flagYaml, flagListDefinedTemplates)
	if err := cmd.Execute(os.Args[1:]...); err != nil {
		fmt.Fprintfln(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

func entrypoint(cmd *cli.Command, args ...string) error {
	if len(args) < 1 {
		return cli.Usage(cmd)
	}

	var t *template.Template
	var err error
	for _, arg := range args {
		t, err = parseFiles(t, arg)
		if err != nil {
			return err
		}
	}

	if flagListDefinedTemplates.Value {
		fmt.Println(t.DefinedTemplates())
		return nil
	}

	data := map[interface{}]interface{}{}
	for _, fileName := range flagYaml.Value {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(data, &data); err != nil {
			return errors.New(fileName + ": " + err.Error())
		}
	}

	if err := t.ExecuteTemplate(os.Stdout, filepath.Base(args[0]), data); err != nil {
		return err
	}
	return nil
}

func parseFiles(t *template.Template, filename string) (*template.Template, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return template.New(filename).Parse(string(b))
	} else if _, err := t.New(filename).Parse(string(b)); err != nil {
		return nil, err
	}
	return t, nil
}
