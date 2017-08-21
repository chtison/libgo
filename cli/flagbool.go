package cli

import "strconv"

// FlagBool ...
type FlagBool struct {
	*flag
	Value bool
}

// NewFlagBool ...
func NewFlagBool(shortName rune, longName string, value bool) *FlagBool {
	return &FlagBool{
		flag:  newFlag(shortName, longName),
		Value: value,
	}
}

// Type ...
func (flag *FlagBool) Type() string {
	return ""
}

// Parse ...
func (flag *FlagBool) Parse(arg *string) error {
	if arg != nil {
		b, err := strconv.ParseBool(*arg)
		if err != nil {
			return err
		}
		flag.Value = b
	} else {
		flag.Value = true
	}
	return nil
}
