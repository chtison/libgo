package tmpl

import (
	"reflect"
	"text/template"

	"github.com/chtison/libgo/tmpl/generated"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"fmt":     generated.NewFmt,
		"json":    generated.NewJson,
		"strings": generated.NewStrings,
		"time":    generated.NewTime,
		"yaml":    generated.NewYaml,
		"isType":  isType,
	}
}

func isType(i interface{}, typeString string) bool {
	if t := reflect.TypeOf(i); t != nil {
		return t.String() == typeString
	}
	return false
}
