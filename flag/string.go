package flag

import "strings"

// String ...
type String struct {
	shortName string
	longName  string
	value     string
	usage     string
}

// NewString ...
func NewString(shortName, longName, defaultValue string) *String {
	if shortName == "" && longName == "" {
		panic("String can't have empty names.")
	}
	if len([]rune(shortName)) > 1 {
		panic("String's short name can't have more than one rune.")
	}
	if strings.Contains(longName, "=") {
		panic("String's long name cannot contains '='")
	}
	return &String{
		shortName: shortName,
		longName:  longName,
		value:     defaultValue,
	}
}

// Name ...
func (str *String) Name() (shortName, longName string) {
	return str.shortName, str.longName
}

// Usage ...
func (str *String) Usage() string {
	return str.usage
}

// SetUsage ...
func (str *String) SetUsage(usage string) {
	str.usage = usage
}

// Value ...
func (str *String) Value() string {
	return str.value
}

// ParseValue ...
func (str *String) ParseValue(value string) error {
	str.value = value
	return nil
}
