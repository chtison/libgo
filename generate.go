// +build ignore

package main

import (
	"go/importer"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chtison/libgo/fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintfln(os.Stderr, "Usage: %s PKGPATH", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if err := generate(os.Args[1]); err != nil {
		fmt.Fprintfln(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

func generate(importPath string) error {
	pkg, err := importer.Default().Import(importPath)
	if err != nil {
		return err
	}
	b := fmt.NewBuilder()
	b.Printfln("package %s", pkg.Name())
	b.Printfln(`import "%s"`, importPath)
	consts := make([]types.Object, 0)
	vars   := make([]types.Object, 0)
	funcs := make([]types.Object, 0)
	typeNames := make([]types.Object, 0)
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		object := scope.Lookup(name)
		if object.Exported() {
			switch object.(type) {
			case *types.Func:
				funcs = append(funcs, object)
			case *types.Const:
				consts = append(consts, object)
			case *types.Var:
				vars = append(vars, object)
			case *types.TypeName:
				typeNames = append(typeNames, object)
			}
		}
	}
	generateObjects(b, pkg.Name(), "const", consts)
	generateObjects(b, pkg.Name(), "var", vars)
	generateObjects(b, pkg.Name(), "var", funcs)
	generateObjects(b, pkg.Name(), "type", typeNames)
	path := filepath.Join(pkg.Name(), "generated.go")
	if err := ioutil.WriteFile(path, []byte(b.String()), 0644); err != nil {
		return err
	}
	fmt.Printfln("+ %s successfully generated", path)
	return nil
}

func generateObjects(b *fmt.Builder, pkgName, category string, objects []types.Object) {
	if len(objects) > 0 {
		b.Printfln("%s (", category)
		for _, object := range objects {
			b.Printfln("    %s = %s.%[1]s", object.Name(), pkgName)
		}
		b.Printfln(")")
	}
}
