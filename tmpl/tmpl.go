package tmpl

import (
	"text/template"

	"github.com/chtison/libgo/deepcopy"
	"github.com/chtison/libgo/tmpl/pkg/generated"
	"github.com/chtison/libgo/tmpl/pkg/utils"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"copy":    deepcopy.Copy,
		"fmt":     generated.NewFmt,
		"json":    generated.NewJson,
		"map":     utils.NewMap,
		"strings": generated.NewStrings,
		"time":    generated.NewTime,
		"istype":  utils.IsType,
		"yaml":    generated.NewYaml,
	}
}
