package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/chtison/libgo/fmt"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printfln(`usage: %s OUTPUT_FILE DATA_YAML TEMPLATE [...]`, os.Args[0])
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		"dictStringToInterface":    dictStringToInterface,
		"dictInterfaceToInterface": dictInterfaceToInterface,
		"time":    NewTime,
		"strings": NewString,
	}
	tmpl, err := template.New(os.Args[3]).Funcs(funcMap).ParseFiles(os.Args[3:]...)
	if err != nil {
		fmt.Fprintfln(os.Stderr, `error: %v: %s`, os.Args[3:], err)
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
	outfile, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintfln(os.Stderr, `error: %s: %s`, os.Args[1], err)
		os.Exit(1)
	}
	if err := tmpl.ExecuteTemplate(outfile, filepath.Base(os.Args[3]), m); err != nil {
		fmt.Fprintfln(os.Stderr, `error: processing template failed: %s`, err)
		outfile.Close()
		os.Remove(os.Args[3])
		os.Exit(1)
	}
	outfile.Close()
}

func dictStringToInterface(d map[string]interface{}, values ...interface{}) map[string]interface{} {
	if d == nil {
		d = make(map[string]interface{})
	}
	for i := 0; i < len(values); i += 2 {
		d[values[i].(string)] = values[i+1]
	}
	return d
}

func dictInterfaceToInterface(d map[interface{}]interface{}, values ...interface{}) map[interface{}]interface{} {
	if d == nil {
		d = make(map[interface{}]interface{})
	}
	for i := 0; i < len(values); i += 2 {
		d[values[i]] = values[i+1]
	}
	return d
}

// Time ...
type Time struct{}

// NewTime ...
func NewTime() *Time {
	return &Time{}
}

// Now ...
func (t *Time) Now() string {
	return time.Now().String()
}

// String ...
type String struct{}

// NewString ...
func NewString() *String {
	return &String{}
}

// Replace ...
func (str *String) Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Split ...
func (str *String) Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// NewReplacer ...
func (str *String) NewReplacer(oldnew ...string) *strings.Replacer {
	return strings.NewReplacer(oldnew...)
}
