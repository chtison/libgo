package shellcolors

import (
	"bytes"
	"fmt"
)

// ShellColor ...
type ShellColor struct {
	codes []CodeSGR
	// temporary storage for computed string in String() method
	str *string
}

// New ...
func New(codes ...CodeSGR) *ShellColor {
	new := &ShellColor{}
	new.Add(codes...)
	return new
}

// NewWithColor ...
func NewWithColor(color uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.Color(color)
	return new
}

// NewWithColorRGB ...
func NewWithColorRGB(red, green, blue uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.ColorRGB(red, green, blue)
	return new
}

// NewWithBgColor ...
func NewWithBgColor(color uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.BgColor(color)
	return new
}

// NewWithBgColorRGB ...
func NewWithBgColorRGB(red, green, blue uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.BgColorRGB(red, green, blue)
	return new
}

// Add adds SGR codes in order to the format string
func (sc *ShellColor) Add(codes ...CodeSGR) *ShellColor {
	sc.str = nil
	for _, code := range codes {
		sc.codes = append(sc.codes, code)
	}
	return sc
}

// Color ...
func (sc *ShellColor) Color(color uint8) *ShellColor {
	sc.Add(CustomColor, CodeSGR(5), CodeSGR(color))
	return sc
}

// ColorRGB ...
func (sc *ShellColor) ColorRGB(red, green, blue uint8) *ShellColor {
	sc.Add(CustomColor, CodeSGR(2),
		CodeSGR(red), CodeSGR(green), CodeSGR(blue))
	return sc
}

// BgColor ...
func (sc *ShellColor) BgColor(color uint8) *ShellColor {
	sc.Add(BgCustomColor, CodeSGR(5), CodeSGR(color))
	return sc
}

// BgColorRGB ...
func (sc *ShellColor) BgColorRGB(red, green, blue uint8) *ShellColor {
	sc.Add(BgCustomColor, CodeSGR(2),
		CodeSGR(red), CodeSGR(green), CodeSGR(blue))
	return sc
}

// String ...
func (sc *ShellColor) String() string {
	if len(sc.codes) == 0 {
		return ""
	}
	if sc.str != nil {
		return *sc.str
	}
	var buf bytes.Buffer
	buf.WriteString(codeSgrStart)
	isFirst := true
	for _, code := range sc.codes {
		if isFirst {
			buf.WriteString(fmt.Sprintf("%d", code))
			isFirst = false
		} else {
			buf.WriteString(fmt.Sprintf("%s%d", codeSgrSeparator, code))
		}
	}
	buf.WriteString(codeSgrEnd)
	str := buf.String()
	sc.str = &str
	return str
}
