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
	b.Printfln(`import "%s"`, importPath)
	b.Printfln("type %s struct {}", typeName)
	b.Printfln("func New%s() *%[1]s { return &%[1]s{} }", typeName)
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		object := scope.Lookup(name)
		if object.Exported() {
			switch object.(type) {
			case *types.Func:
				generateSignature(b, object, importPathBase, typeName)
			case *types.Const, *types.Var:
				if _, ok := object.Type().(*types.Signature); ok {
					generateSignature(b, object, importPathBase, typeName)
					break
				}
				b.Printf("func (%s) %s() ", typeName, object.Name())
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
	p := filepath.Join("generated", importPathBase+".go")
	if err := ioutil.WriteFile(p, []byte(b.String()), 0644); err != nil {
		return err
	}
	fmt.Printfln("+ %s successfully generated", p)
	return nil
}

func generateSignature(b *fmt.Builder, object types.Object, importPathBase, typeName string) {
	sig := object.Type().(*types.Signature)
	b.Println("")
	b.Printf("func (*%s) %s(", typeName, object.Name())
	generateSignatureParams(b, sig, true)
	b.Print(") ")
	results := sig.Results()
	if results.Len() > 1 {
		b.Print("(")
	}
	for i := 0; i < results.Len(); i++ {
		if i > 0 {
			b.Print(", ")
		}
		b.Print(types.TypeString(results.At(i).Type(), qualifier))
	}
	if results.Len() > 1 {
		b.Print(")")
	}
	b.Println(" {")
	if results != nil && results.Len() > 0 {
		b.Print("return ")
	}
	b.Printf("%s.%s(", importPathBase, object.Name())
	generateSignatureParams(b, sig, false)
	if sig.Variadic() {
		b.Print("...")
	}
	b.Println(")")
	b.Println("}")
}

func generateSignatureParams(b *fmt.Builder, sig *types.Signature, printType bool) {
	params := sig.Params()
	if params == nil {
		return
	}
	for i := 0; i < params.Len(); i++ {
		if i > 0 {
			b.Print(", ")
		}
		b.Print(params.At(i).Name())
		if printType {
			s := types.TypeString(params.At(i).Type(), qualifier)
			if i < params.Len()-1 || !sig.Variadic() {
				b.Printf(" %s", s)
			} else {
				s = strings.Replace(s, "[]", "...", 1)
				b.Printf(" %s", s)
			}
		}
	}
}

func qualifier(p *types.Package) string {
	return path.Base(p.Name())
}
