// +build ignore

package main

import (
	"go/importer"
	"go/types"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

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
	importPathBase := filepath.Base(importPath)
	typeName := strings.Title(importPathBase)
	b.Println("package generated")
	b.Printfln("type %s struct {}", typeName)
	b.Printfln("func New%s() *%[1]s { return &%[1]s{} }", typeName)
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		object := scope.Lookup(name)
		if object.Exported() {
			switch object.(type) {
			case *types.Func:
				t := object.Type().(*types.Signature)
				b.Println("")
				b.Printf("func (_ *%s) ", typeName)
				b.Print(strings.SplitN(types.ObjectString(object, func(pkg *types.Package) string { return path.Base(pkg.Path()) }), ".", 2)[1])
				b.Println(" {")
				if t.Results() != nil {
					b.Print("return ")
				}
				b.Printf("%s.%s(", importPathBase, object.Name())
				params := t.Params()
				for i := 0; i < params.Len(); i++ {
					if i > 0 {
						b.Print(", ")
					}
					b.Printf("%s", params.At(i).Name())
				}
				if t.Variadic() {
					b.Print("...")
				}
				b.Println(")")
				b.Println("}")
			case *types.Const, *types.Var:
				b.Printf("func (_ %s) %s() ", typeName, object.Name())
				switch t := object.Type().(type) {
				case *types.Basic:
					b.Printf("%s", strings.TrimPrefix(t.String(), "untyped "))
				default:
					b.Printf("%s", t)
				}
				b.Printfln(" { return %s.%s }", object.Pkg().Name(), object.Name())
			}
		}
	}
	os.MkdirAll("generated", 0755)
	p := filepath.Join("generated", importPathBase + ".go")
	if err := ioutil.WriteFile(p, []byte(b.String()), 0644); err != nil {
		return err
	}
	fmt.Printfln("+ %s successfully generated", p)
	return nil
}
