package cli

// FlagStringList ...
type FlagStringList struct {
	*flag
	Value []string
}

// NewFlagStringList ...
func NewFlagStringList(shortName rune, longName string, value []string) *FlagStringList {
	return &FlagStringList{
		flag:  newFlag(shortName, longName),
		Value: value,
	}
}

// Type ...
func (flag *FlagStringList) Type() string {
	return "[]string"
}

// Parse ...
func (flag *FlagStringList) Parse(arg *string) error {
	if arg == nil {
		return &ErrFlagNeedsArg{NbrOfArgs: 1}
	}
	if flag.Value == nil {
		flag.Value = make([]string, 0, 1)
	}
	flag.Value = append(flag.Value, *arg)
	return nil
}
