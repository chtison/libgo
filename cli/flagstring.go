package cli

// FlagString ...
type FlagString struct {
	*flag
	Value string
}

// NewFlagString ...
func NewFlagString(shortName rune, longName string, value string) *FlagString {
	return &FlagString{
		flag:  newFlag(shortName, longName),
		Value: value,
	}
}

// Type ...
func (flag *FlagString) Type() string {
	return "string"
}

// Parse ...
func (flag *FlagString) Parse(arg *string) error {
	if arg == nil {
		return &ErrFlagNeedsArg{NbrOfArgs: 1}
	}
	flag.Value = *arg
	return nil
}
