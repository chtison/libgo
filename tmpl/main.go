package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/chtison/libgo/fmt"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printfln(`usage: %s TEMPLATE YAMLFILE RESULT [TEMPLATE, [...]]`, os.Args[0])
		os.Exit(1)
	}
	funcMap := template.FuncMap{}
	tmpl, err := template.New(os.Args[1]).Funcs(funcMap).ParseFiles(os.Args[1])
	if err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[1], err)
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[2], err)
		os.Exit(1)
	}
	var m map[string]interface{}
	if err = yaml.Unmarshal(data, &m); err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[2], err)
		os.Exit(1)
	}
	outfile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[3], err)
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		if _, err = tmpl.ParseFiles(os.Args[4:]...); err != nil {
			fmt.Fprintfln(os.Stderr, `error: %s`, err)
			outfile.Close()
			os.Remove(os.Args[3])
			os.Exit(1)
		}
	}
	if tmpl.Execute(outfile, m); err != nil {
		fmt.Fprintfln(os.Stderr, `error: processing template failed: %s`, err)
		outfile.Close()
		os.Remove(os.Args[3])
		os.Exit(1)
	}
	outfile.Close()
}
