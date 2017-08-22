package cli

import "strings"

type (
	// Command ...
	Command struct {
		name     string
		Function func(*Command, ...string) error
		Flags    *FlagSet
		Children CommandSet
		parent   *Command
		Usage    CommandUsage
	}
	// CommandUsage ...
	CommandUsage struct {
		Arguments string
		Synopsys  string
	}
	// CommandSet ...
	CommandSet map[string]*Command
)

////
//////// Command
////

// NewCommand ...
func NewCommand(name string) *Command {
	if name == "" {
		panic("A command can't have an empty name")
	}
	if strings.HasPrefix(name, "-") {
		panic("A command can't have a name beginning with '-'")
	}
	cmd := &Command{
		name:     name,
		Flags:    NewFlagSet(),
		Children: NewCommandSet(),
		Function: Usage,
	}
	flagHelp := NewFlagBool(0, "help", false)
	flagHelp.Usage().Synopsys = "Print usage"
	cmd.Flags.Add(flagHelp)
	return cmd
}

// Name ...
func (cmd *Command) Name() string {
	return cmd.name
}

// FindFlag ...
func (cmd *Command) FindFlag(shortName rune, longName string) Flag {
	for p := cmd; p != nil; p = cmd.parent {
		if flags := p.Flags; flags != nil {
			if flag := flags.Find(shortName, longName); flag != nil {
				return flag
			}
		}
	}
	return nil
}

// Execute ...
func (cmd *Command) Execute(commandLine ...string) error {
	return cmd.execute(false, commandLine...)
}

func (cmd *Command) execute(flagStop bool, commandLine ...string) error {
	i := 0
	for ; i < len(commandLine); i++ {
		arg := []rune(commandLine[i])
		if !flagStop && len(arg) > 1 && arg[0] == '-' {
			// Handle long flag
			if arg[1] == '-' {
				if len(arg) == 2 {
					flagStop = true
					continue
				}
				if flag := cmd.FindFlag(0, string(arg[2:])); flag != nil {
					var flagParseArg *string
					for {
						if err := flag.Parse(flagParseArg); err != nil {
							if err, ok := err.(*ErrFlagNeedsArg); ok {
								i++
								if i >= len(commandLine) {
									return &ErrFlagNeedsArg{string(arg), flag, err.NbrOfArgs}
								}
								flagParseArg = &commandLine[i]
								continue
							}
							return &ErrFlagParseError{string(arg), flag, err}
						}
						break
					}
					continue
				}
				return &ErrFlagNotFound{string(arg)}
			}
			// Handler short flags
			for j := 1; j < len(arg); j++ {
				shortName := arg[j]
				if flag := cmd.FindFlag(shortName, ""); flag != nil {
					var flagParseArg *string
					for k := j; true; {
						if err := flag.Parse(flagParseArg); err != nil {
							if err, ok := err.(*ErrFlagNeedsArg); ok {
								if j+1 < len(arg) {
									s := string(arg[j+1:])
									flagParseArg = &s
									j = len(arg)
								} else {
									i++
									if i >= len(commandLine) {
										return &ErrFlagNeedsArg{"-" + string(shortName), flag, err.NbrOfArgs}
									}
									flagParseArg = &commandLine[i]
								}
								continue
							}
							return &ErrFlagParseError{"-" + string(arg[k]), flag, err}
						}
						break
					}
					continue
				}
				return &ErrFlagNotFound{"-" + string(shortName)}
			}
			continue
		}
		// Handle child command
		if cmd.Children != nil {
			if child := cmd.Children.Find(string(arg)); child != nil {
				child.parent = cmd
				return child.execute(flagStop, commandLine[i+1:]...)
			}
		}
		break
	}
	// Handle --help flag
	if cmd.Flags != nil {
		if flag := cmd.Flags.Find(0, "help"); flag != nil {
			if flag, ok := flag.(*FlagBool); ok {
				if flag.Value {
					return Usage(cmd, commandLine[i:]...)
				}
			}
		}
	}
	// Handle self execution
	if cmd.Function != nil {
		if err := cmd.Function(cmd, commandLine[i:]...); err != nil {
			return &ErrCommandFunction{err}
		}
		return nil
	}
	return nil
}

////
//////// CommandSet
////

// NewCommandSet ...
func NewCommandSet() CommandSet {
	return make(map[string]*Command)
}

// Find ...
func (set CommandSet) Find(name string) *Command {
	return set[name]
}

// Add ...
func (set CommandSet) Add(commands ...*Command) {
	for _, cmd := range commands {
		set[cmd.name] = cmd
	}
}
