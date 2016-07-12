# Golang package baseconverter

[![GoDoc](https://godoc.org/github.com/chtison/libgo/baseconverter?status.svg)](https://godoc.org/github.com/chtison/libgo/baseconverter)
[![Build Status](https://travis-ci.org/chtison/libgo.svg?branch=master)](https://travis-ci.org/chtison/libgo)

Package baseconverter is a set of functions which perform numerical base conversion.

### Install this package
```
$ go get -v github.com/chtison/libgo/baseconverter/...
```

### Short documentation

A number is represented as a uint in decimal base or as a string (interpreted
as UTF-8 encoded) in any base.
``` go
        var number uint    // decimal base (base 10)
        var number string  // any base, even decimal one
```

A base is represented as a string (interpreted as UTF-8 encoded), and must own
at least two different runes.
```go
        var base string
        len([]rune(base)) >= 2
        base[i] != base[j] with i != j
```

For example, you can convert a decimal number to base 16:
```go
        package main

        import (
                "fmt"

                bc "github.com/chtison/libgo/baseconverter"
        )

        func main() {
                nbrInBase16, _ := bc.DecimalToBase(51966, "0123456789abcdef")
                fmt.Println(nbrInBase16)
        }
```
Or convert back a number in base "01" (base 2) to base 10:
```go
        package main

        import (
                "fmt"

                bc "github.com/chtison/libgo/baseconverter"
        )

        func main() {
                nbr, _ := bc.BaseToDecimal("101010", "01")
                fmt.Println(nbr)
        }
```
