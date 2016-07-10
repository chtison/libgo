package shellcolors

import (
	"fmt"
	"strings"
	"testing"
)

func printError(t *testing.T, p1, p2, p3 string) {
	p2 = strings.Replace(p2, "\033", "\\033", -1)
	p3 = strings.Replace(p3, "\033", "\\033", -1)
	t.Errorf("%s == \"%s\". Expected: \"%s\"\n", p1, p2, p3)
}

// *****************************************************************************
// Test New()
func TestNew(t *testing.T) {
	testNew(t, "New().String()", "")
	testNew(t, "New(Bold, Underline).String()", "\033[1;4m", Bold, Underline)
	testNew(t, "New(Reset).String()", "\033[0m", Reset)
}

func testNew(t *testing.T, functionCall, expected string, params ...CodeSGR) {
	new := New(params...)
	if str := new.String(); str != expected {
		printError(t, functionCall, str, expected)
	}
}

// *****************************************************************************
// Test Color()
func TestColor(t *testing.T) {
	for i := uint16(0); i < 256; i++ {
		testColor(t, fmt.Sprintf("New().Color(%d)", i),
			fmt.Sprintf("\033[38;5;%dm", i), uint8(i))
	}
}

func testColor(t *testing.T, functionCall, expected string, color uint8) {
	new := New().Color(color)
	if str := new.String(); str != expected {
		printError(t, functionCall, str, expected)
	}
}

// *****************************************************************************
// Test ColorRGB()
func TestColorRGB(t *testing.T) {
	var r, g, b uint16
	for ; r < 256; r++ {
		for ; g < 256; g++ {
			for ; b < 256; b++ {
				testColorRGB(t,
					fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b),
					uint8(r), uint8(g), uint8(b))
			}
		}
	}
}

func testColorRGB(t *testing.T, expected string, r, g, b uint8) {
	new := New().ColorRGB(r, g, b)
	if str := new.String(); str != expected {
		printError(t, fmt.Sprintf("New().ColorRGB(%d, %d, %d)", r, g, b),
			str, expected)
	}
}
