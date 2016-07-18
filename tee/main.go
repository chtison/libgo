package main

// TODO(): ignore SIGINT when -i flag set

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const usage = `usage: tee [-a] [filename [...]]
tee writes what it reads from standard input to standard output and files
`

func printUsage() {
	fmt.Fprint(os.Stderr, usage)
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	flagAppend = flag.Bool("a", false, "`append` the output to the files rather than truncating them")
)

func main() {

	flag.Usage = printUsage
	flag.Parse()

	os.Exit(tee())
}

func tee() (exit int) {

	files := make([]io.WriteCloser, 0, flag.NArg())
	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	flags := os.O_WRONLY | os.O_CREATE
	if *flagAppend {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	for _, arg := range flag.Args() {
		fd, err := os.OpenFile(arg, flags, 0666)
		if err != nil {
			log.Println(err)
			exit = 1
			continue
		}
		files = append(files, fd)
	}

	writers := make([]io.Writer, len(files))
	for i, file := range files {
		writers[i] = io.Writer(file)
	}
	multiWriter := io.MultiWriter(writers...)

	teeReader := io.TeeReader(os.Stdin, multiWriter)

	_, err := io.Copy(os.Stdout, teeReader)
	if err != nil {
		log.Println(err)
		exit = 1
	}
	return
}
