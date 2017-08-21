package cli

import "fmt"

type (
	// ErrFlagNotFound ...
	ErrFlagNotFound struct {
		FlagName string
	}
	// ErrFlagParseError ...
	ErrFlagParseError struct {
		FlagName   string
		Flag       Flag
		ParseError error
	}
	// ErrFlagNeedsArg ...
	ErrFlagNeedsArg struct {
		FlagName  string
		Flag      Flag
		NbrOfArgs int
	}
	// ErrCommandFunction ...
	ErrCommandFunction struct {
		error
	}
)

func (err *ErrFlagNotFound) Error() string {
	return fmt.Sprintf(`flag not found: "%s"`, err.FlagName)
}

func (err *ErrFlagNeedsArg) Error() string {
	if err.NbrOfArgs < 2 {
		return fmt.Sprintf(`%s: flag needs 1 additional argument`, err.FlagName)
	}
	return fmt.Sprintf(`%s: flag needs %d additionals arguments`, err.FlagName, err.NbrOfArgs)
}

func (err *ErrFlagParseError) Error() string {
	return fmt.Sprintf(`%s: %s`, err.FlagName, err.ParseError)
}
