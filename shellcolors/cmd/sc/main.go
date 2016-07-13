/*
Package sc is a command line interface for the package shellcolors.

	$ sc green bold ; echo "Hello World !" ; sc reset
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	sc "github.com/chtison/libgo/shellcolors"
)

type argParser func(string, *sc.ShellColor) (bool, error)
type argParserConstructor func() argParser

type paramToCodeSGR struct {
	param                string
	codeSGR              sc.CodeSGR
	argParserConstructor argParserConstructor
	usage                string
}

var tabParamToCodeSGR = [...]paramToCodeSGR{
	{"reset", sc.Reset, nil, ""},
	{"bold", sc.Bold, nil, ""},
	{"faint", sc.Faint, nil, ""},
	{"italic", sc.Italic, nil, ""},
	{"underline", sc.Underline, nil, ""},
	{"blinkslow", sc.BlinkSlow, nil, ""},
	{"blinkrapid", sc.BlinkRapid, nil, ""},
	{"negative", sc.Negative, nil, ""},
	{"conceal", sc.Conceal, nil, ""},
	{"crossedout", sc.CrossedOut, nil, ""},

	{"fraktur", sc.Fraktur, nil, ""},
	{"noBold", sc.NoBold, nil, ""},
	{"noBoldAndFaint", sc.NoBoldAndFaint, nil, ""},
	{"noItalicAndFraktur", sc.NoItalicAndFraktur, nil, ""},
	{"noUnderline", sc.NoUnderline, nil, ""},
	{"noBlink", sc.NoBlink, nil, ""},
	{"reserved26", sc.Reserved26, nil, ""},
	{"noNegative", sc.NoNegative, nil, ""},
	{"noConceal", sc.NoConceal, nil, ""},
	{"noCrossedOut", sc.NoCrossedOut, nil, ""},

	{"black", sc.Black, nil, ""},
	{"red", sc.Red, nil, ""},
	{"green", sc.Green, nil, ""},
	{"yellow", sc.Yellow, nil, ""},
	{"blue", sc.Blue, nil, ""},
	{"magenta", sc.Magenta, nil, ""},
	{"cyan", sc.Cyan, nil, ""},
	{"white", sc.White, nil, ""},
	{"color", sc.CustomColor, customColor, "{0-255}"},
	{"colorRGB", sc.CustomColor, customColorRGB, "{0-255} {0-255} {0-255}"},
	{"default", sc.DefaultColor, nil, ""},

	{"bgBlack", sc.BgBlack, nil, ""},
	{"bgRed", sc.BgRed, nil, ""},
	{"bgGreen", sc.BgGreen, nil, ""},
	{"bgYellow", sc.BgYellow, nil, ""},
	{"bgBlue", sc.BgBlue, nil, ""},
	{"bgMagenta", sc.BgMagenta, nil, ""},
	{"bgCyan", sc.BgCyan, nil, ""},
	{"bgWhite", sc.BgWhite, nil, ""},
	{"bgcolor", sc.BgCustomColor, customBgColor, "{0-255}"},
	{"bgcolorRGB", sc.BgCustomColor, customBgColorRGB, "{0-255} {0-255} {0-255}"},
	{"bgDefaultColor", sc.BgDefaultColor, nil, ""},

	{"blackHI", sc.BlackHI, nil, ""},
	{"redHI", sc.RedHI, nil, ""},
	{"greenHI", sc.GreenHI, nil, ""},
	{"yellowHI", sc.YellowHI, nil, ""},
	{"blueHI", sc.BlueHI, nil, ""},
	{"magentaHI", sc.MagentaHI, nil, ""},
	{"cyanHI", sc.CyanHI, nil, ""},
	{"whiteHI", sc.WhiteHI, nil, ""},

	{"bgBlackHI", sc.BgBlackHI, nil, ""},
	{"bgRedHI", sc.BgRedHI, nil, ""},
	{"bgGreenHI", sc.BgGreenHI, nil, ""},
	{"bgYellowHI", sc.BgYellowHI, nil, ""},
	{"bgBlueHI", sc.BgBlueHI, nil, ""},
	{"bgMagentaHI", sc.BgMagentaHI, nil, ""},
	{"bgCyanHI", sc.BgCyanHI, nil, ""},
	{"bgWhiteHI", sc.BgWhiteHI, nil, ""},
}

func custom(nbOfArgs uint8, finish func([]uint8, *sc.ShellColor)) argParser {
	var Err error
	a := make([]uint8, 0, nbOfArgs)
	return func(arg string, shellColor *sc.ShellColor) (bool, error) {
		color, err := strconv.ParseUint(arg, 10, 8)
		if err != nil {
			if Err == nil {
				Err = fmt.Errorf("%s", arg)
			}
		} else {
			a = append(a, uint8(color))
		}
		if len(a) < cap(a) {
			return false, nil
		}
		finish(a, shellColor)
		return true, Err
	}
}

func customColor() argParser {
	return custom(1, func(args []uint8, shellColor *sc.ShellColor) {
		shellColor.Color(args[0])
	})
}

func customColorRGB() argParser {
	return custom(3, func(args []uint8, shellColor *sc.ShellColor) {
		fmt.Println(args)
		shellColor.ColorRGB(args[0], args[1], args[2])
	})
}

func customBgColor() argParser {
	return custom(1, func(args []uint8, shellColor *sc.ShellColor) {
		shellColor.BgColor(args[0])
	})
}

func customBgColorRGB() argParser {
	return custom(3, func(args []uint8, shellColor *sc.ShellColor) {
		shellColor.BgColorRGB(args[0], args[1], args[2])
	})
}

func searchCodeSGR(param string) *paramToCodeSGR {
	for i := range tabParamToCodeSGR {
		if tabParamToCodeSGR[i].param == param {
			return &tabParamToCodeSGR[i]
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "list" {
			printCommands()
			return
		} else if os.Args[1] == "examples" {
			printExample()
			return
		}
	}
	shellColor := sc.New()
	var ap argParser
	for _, arg := range os.Args[1:] {
		if ap != nil {
			end, err := ap(arg, shellColor)
			if err != nil {
				printError(err.Error())
				ap = nil
			}
			if end {
				ap = nil
			}
		} else if paramToCode := searchCodeSGR(arg); paramToCode != nil {
			if paramToCode.argParserConstructor != nil {
				ap = paramToCode.argParserConstructor()
			} else {
				shellColor.Add(paramToCode.codeSGR)
			}
		} else {
			printError(arg)
		}
	}
	fmt.Print(shellColor)
}

const usage = `usage: sc {cmd}+
       sc list

Shellcolors lets you modify style of you outputed text in a terminal
like the boldness or the background color.
However, some commands might not be supported by your terminal.

Type 'sc list' to get all available commands that sc accept (case sensitive).

Example: sc bold green ; echo "Hello World !" ; sc reset
`

func printUsage() {
	fmt.Print(usage)
}

func printCommands() {
	reset := sc.New(sc.Reset)
	for _, paramToCode := range tabParamToCodeSGR {
		if paramToCode.usage != "" {
			fmt.Println(fmt.Sprintf("- %s %s",
				paramToCode.param,
				paramToCode.usage))
		} else {
			fmt.Println(fmt.Sprintf("- %s%s%s",
				sc.New(paramToCode.codeSGR),
				paramToCode.param,
				reset))
		}
	}
}

func printExample() {
	fmt.Println(fmt.Sprintf(
		"sc color 42 bold underline ; echo \"Hello World 1\" ; sc reset\n%sHello World !%s",
		sc.NewWithColor(42, sc.Bold, sc.Underline),
		sc.New(sc.Reset),
	))
	fmt.Println(fmt.Sprintf(
		"sc colorRGB 255 153 51 bold ; echo \"Hello World 1\" ; sc reset\n%sHello World !%s",
		sc.NewWithColorRGB(255, 153, 51, sc.Bold),
		sc.New(sc.Reset),
	))
}

func printError(arg string) {
	fmt.Println(fmt.Sprintf("%ssc: warning:%s bad parameter \"%s\"",
		sc.NewWithColor(208, sc.Bold),
		sc.New(sc.Reset),
		arg))
}
