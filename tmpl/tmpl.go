package tmpl

//go:generate ./generate.sh

import (
	"reflect"
	"text/template"

	"github.com/chtison/libgo/tmpl/generated"
)

// Funcs ...
func Funcs() template.FuncMap {
	return template.FuncMap{
		"fmt":     generated.NewFmt,
		"strings": generated.NewStrings,
		"time":    generated.NewTime,
		"isType":  isType,
	}
}

func isType(i interface{}, typeString string) bool {
	if t := reflect.TypeOf(i); t != nil {
		return t.String() == typeString
	}
	return false
}
