package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/chtison/libgo/fmt"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printfln(`usage: %s TEMPLATE JSONFILE RESULT`, os.Args[0])
		os.Exit(1)
	}
	tmpl, err := template.ParseFiles(os.Args[1])
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
	if err = json.Unmarshal(data, &m); err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[2], err)
		os.Exit(1)
	}
	outfile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[3], err)
		os.Exit(1)
	}
	defer outfile.Close()
	if tmpl.Execute(outfile, m); err != nil {
		fmt.Fprintfln(os.Stderr, `error: processing template failed: %s`, err)
		outfile.Close()
		os.Remove(os.Args[3])
		os.Exit(1)
	}
}
