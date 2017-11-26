package tmpl

//go:generate ./generate.sh

import (
	"text/template"

	"github.com/chtison/libgo/tmpl/generated"
)

// Funcs ...
func Funcs() template.FuncMap {
	return template.FuncMap{
		"fmt":     generated.NewFmt,
		"strings": generated.NewStrings,
		"time":    generated.NewTime,
	}
}
