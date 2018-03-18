package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/chtison/libgo/tmpl"
	"github.com/chtison/libgo/yaml"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Version: "1.10.0",
		Use: fmt.Sprintf("%s [FLAGS] TEMPLATE [TEMPLATE ...]", filepath.Base(os.Args[0])),
		Short: "A CLI for the golang template engine",
		DisableSuggestions:    true,
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Args:  cobra.MinimumNArgs(1),
		RunE: handler,
	}
	flagYamlData []string
)

func main() {
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().StringArrayVarP(&flagYamlData, "yaml", "y", nil, "Paths to yaml data files")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func handler(cmd *cobra.Command, args []string) error {
	var data interface{}
	for _, f := range flagYamlData {
		if err := yaml.ReadFile(f, &data); err != nil {
			return err
		}
	}
	t := template.New("").Funcs(tmpl.Funcs())
	for _, filename := range args {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		if _, err := t.New(filename).Parse(string(b)); err != nil {
			return err
		}
	}
	if err := t.ExecuteTemplate(os.Stdout, args[0], data); err != nil {
		return err
	}
	return nil
}
