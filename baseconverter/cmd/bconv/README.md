# Golang package bconv

[![GoDoc](https://godoc.org/github.com/chtison/libgo/baseconverter/cmd/bconv?status.svg)](https://godoc.org/github.com/chtison/libgo/baseconverter/cmd/bconv)
[![Build Status](https://travis-ci.org/chtison/libgo.svg?branch=master)](https://travis-ci.org/chtison/libgo)

Package bconv is a command line interface for the package [baseconverter](../..).

## Install
	$ go get -v github.com/chtison/libgo/baseconverter/cmd/bconv
	
## Examples
	$ bconv 51966 0123456789 0123456789abcdef
	cafe
	$ bconv 101010 01 0123456789
	42

## HTTP Frontend
	$ bconv serve ':4242'
	Listen on [::]:4242
