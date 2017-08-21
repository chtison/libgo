package tmpl

import (
	"strings"
	"text/template"
	"time"
)

// Funcs ...
func Funcs() template.FuncMap {
	return template.FuncMap{
		"dictStringToInterface":    dictStringToInterface,
		"dictInterfaceToInterface": dictInterfaceToInterface,
		"time":    NewTime,
		"strings": NewString,
	}
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
