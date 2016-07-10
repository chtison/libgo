package main

import (
	"fmt"
	sc "github.com/chtison/libgo/shellcolors"
	"os"
)

var paramToCodeSGR = [...]struct {
	param   string
	codeSGR sc.CodeSGR
}{
	{param: "reset", codeSGR: sc.Reset},
	{"bold", sc.Bold},
	{"faint", sc.Faint},
	{"italic", sc.Italic},
	{"underline", sc.Underline},
	{"blinkslow", sc.BlinkSlow},
	{"blinkrapid", sc.BlinkRapid},
	{"negative", sc.Negative},
	{"conceal", sc.Conceal},
	{"crossedout", sc.CrossedOut},

	{"black", sc.Black},
	{"red", sc.Red},
	{"green", sc.Green},
	{"yellow", sc.Yellow},
	{"blue", sc.Blue},
	{"magenta", sc.Magenta},
	{"cyan", sc.Cyan},
	{"white", sc.White},
	{"custom", sc.CustomColor},
	{"default", sc.DefaultColor},
}

func searchCodeSGR(param string) *sc.CodeSGR {
	for i := range paramToCodeSGR {
		if paramToCodeSGR[i].param == param {
			return &paramToCodeSGR[i].codeSGR
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	shellColor := sc.New()
	for _, arg := range os.Args[1:] {
		if codeSGR := searchCodeSGR(arg); codeSGR != nil {
			shellColor.Add(*codeSGR)
		} else {
			printError(arg)
		}
	}
	fmt.Print(shellColor)
}

const usage = `%s>>> %sUsage: %ssc cmd [, cmd]*
%s- %sShellColor%s writes codes that can be understand by terminals in order to
   modify some parameters like the boldness or the color of the text.
%s- %sExample:%s sc bold green ; echo "Hello World !" ; sc reset
%s- %sDocumentation:%s https://en.wikipedia.org/wiki/ANSI_escape_code#graphics
%s- %sCommands:%s
`

func printUsage() {
	fmt.Printf(usage,
		sc.New(sc.Bold, sc.Red),
		sc.New(sc.Green),
		sc.New(sc.Blue),
		//
		sc.New(sc.Red),
		sc.New(sc.Green),
		sc.New(sc.Reset),
		//
		sc.New(sc.Bold, sc.Red),
		sc.New(sc.Green),
		sc.New(sc.Reset),
		//
		sc.New(sc.Bold, sc.Red),
		sc.New(sc.Green),
		sc.New(sc.Reset),
		//
		sc.New(sc.Bold, sc.Red),
		sc.New(sc.Green),
		sc.New(sc.Reset))

	reset := sc.New(sc.Reset)
	for i := range paramToCodeSGR {
		fmt.Println(fmt.Sprintf("%s%s%s",
			sc.New(paramToCodeSGR[i].codeSGR),
			paramToCodeSGR[i].param,
			reset))
	}
}

func printError(arg string) {
	fmt.Println(fmt.Sprintf("%sWarning%s:%s parameter \"%s\" unknown",
		sc.NewWithColor(208, sc.Underline, sc.Bold),
		sc.New(sc.NoUnderline),
		sc.New(sc.Reset),
		arg))
}
