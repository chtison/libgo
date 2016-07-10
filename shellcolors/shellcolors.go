package shellcolors

import (
	"bytes"
	"fmt"
)

type ShellColor struct {
	codes []CodeSGR
	// temporary storage for computed string in String() method
	str *string
}

func New(codes ...CodeSGR) *ShellColor {
	new := &ShellColor{}
	new.Add(codes...)
	return new
}

func NewWithColor(color uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.Color(color)
	return new
}

func NewWithColorRGB(red, green, blue uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.ColorRGB(red, green, blue)
	return new
}

func NewWithBgColor(color uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.BgColor(color)
	return new
}

func NewWithBgColorRGB(red, green, blue uint8, codes ...CodeSGR) *ShellColor {
	new := New(codes...)
	new.BgColorRGB(red, green, blue)
	return new
}

// Add adds SGR codes in order to the format string
func (self *ShellColor) Add(codes ...CodeSGR) *ShellColor {
	self.str = nil
	for _, code := range codes {
		self.codes = append(self.codes, code)
	}
	return self
}

func (self *ShellColor) Color(color uint8) *ShellColor {
	self.Add(CustomColor, CodeSGR(5), CodeSGR(color))
	return self
}

func (self *ShellColor) ColorRGB(red, green, blue uint8) *ShellColor {
	self.Add(CustomColor, CodeSGR(2),
		CodeSGR(red), CodeSGR(green), CodeSGR(blue))
	return self
}

func (self *ShellColor) BgColor(color uint8) *ShellColor {
	self.Add(BgCustomColor, CodeSGR(5), CodeSGR(color))
	return self
}

func (self *ShellColor) BgColorRGB(red, green, blue uint8) *ShellColor {
	self.Add(BgCustomColor, CodeSGR(2),
		CodeSGR(red), CodeSGR(green), CodeSGR(blue))
	return self
}

func (self *ShellColor) String() string {
	if len(self.codes) == 0 {
		return ""
	}
	if self.str != nil {
		return *self.str
	}
	var buf bytes.Buffer
	buf.WriteString(codeSGR_start)
	isFirst := true
	for _, code := range self.codes {
		if isFirst {
			buf.WriteString(fmt.Sprintf("%d", code))
			isFirst = false
		} else {
			buf.WriteString(fmt.Sprintf("%s%d", codeSGR_separator, code))
		}
	}
	buf.WriteString(codeSGR_end)
	str := buf.String()
	self.str = &str
	return str
}
