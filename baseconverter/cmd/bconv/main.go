/*
Package bconv is a command line interface for the package baseconverter.

	$ bconv 51966 0123456789 0123456789abcdef
	cafe
	$ bconv 101010 01 0123456789
	42
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	bc "github.com/chtison/libgo/baseconverter"
	_ "github.com/chtison/libgo/baseconverter/appengine"
)

func main() {

	if len(os.Args) != 4 {
		if (len(os.Args) == 2 || len(os.Args) == 3) && os.Args[1] == "serve" {
			if len(os.Args) == 3 {
				listenAndServe(os.Args[2])
			} else {
				listenAndServe(":8080")
			}
			return
		}
		printUsage()
		return
	}
	number, e1, e2 := bc.BaseToBase(os.Args[1], os.Args[2], os.Args[3])
	if e1 != nil {
		printError(os.Args[2], e1)
		os.Exit(1)
	} else if e2 != nil {
		printError(os.Args[3], e2)
		os.Exit(1)
	}
	fmt.Println(number)
}

func listenAndServe(laddr string) {
	dir := path.Join(os.Getenv("GOPATH"), "src/github.com/chtison/libgo/baseconverter/appengine")
	if err := os.Chdir(dir); err != nil {
		log.Fatalln("GOPATH is invalid or package is missing:", err)
	}
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/static/", handleStatic)
	fmt.Println("Listen on", l.Addr().String())
	log.Fatalln(http.Serve(l, nil))
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	file := path.Join(os.Getenv("GOPATH"), "src/github.com/chtison/libgo/baseconverter/appengine", r.URL.Path)
	http.ServeFile(w, r, file)
	return
}

func printUsage() {
	fmt.Print(usage)
}

const usage = `usage: bconv {number} {inBase} {toBase}
       bconv serve [[addr]:port]

BaseConverter converts number from base inBase to base toBase.

Example: bconv 51966 0123456789 0123456789abcdef

BaseConverter with serve command starts an HTTP application
on [[addr]:port] if specified, or ':8080' otherwise.

Example: bconv serve '127.0.0.1:4242'
`

func printError(base string, err error) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(`error: "%s": %s`, base, err))
}
