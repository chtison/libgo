package flag

import (
	"fmt"
	"os"
	"strings"
)

// Command ...
type Command struct {
	name  string
	usage string

	parent      *Command
	subCommands *CommandSet

	Function    func(command *Command, args []string) error
	LocalFlags  *FlagSet
	GlobalFlags *FlagSet
}

// NewCommand ...
func NewCommand(name string) *Command {
	command := &Command{
		subCommands: NewCommandSet(),
		LocalFlags:  NewFlagSet(),
		GlobalFlags: NewFlagSet(),
	}
	command.SetName(name)
	return command
}

// Name ...
func (command *Command) Name() string {
	return command.name
}

// SetName ...
func (command *Command) SetName(name string) {
	if name == "" {
		panic("Command's name can't be empty.")
	}
	command.name = name
}

// Usage ...
func (command *Command) Usage() string {
	return command.usage
}

// SetUsage ...
func (command *Command) SetUsage(usage string) {
	command.usage = usage
}

// Parent ...
func (command *Command) Parent() *Command {
	return command.parent
}

// FindFlag ...
func (command *Command) FindFlag(shortName, longName string) Flag {
	if flag := command.LocalFlags.Find(shortName, longName); flag != nil {
		return flag
	}
	for command != nil {
		if flag := command.GlobalFlags.Find(shortName, longName); flag != nil {
			return flag
		}
		command = command.parent
	}
	return nil
}

// RemoveSubCommand ...
func (command *Command) RemoveSubCommand(commands ...*Command) {
	for _, cmd := range commands {
		if cmd.parent != command {
			return
		}
		command.subCommands.Remove(cmd)
		cmd.parent = nil
	}
}

// AddSubCommand ...
func (command *Command) AddSubCommand(commands ...*Command) {
	for _, cmd := range commands {
		if cmd.parent != nil {
			cmd.parent.RemoveSubCommand(cmd)
		}
		command.subCommands.Add(cmd)
		cmd.parent = command
	}
}

// FindSubCommmand ...
func (command *Command) FindSubCommmand(name string) *Command {
	return command.subCommands.Find(name)
}

// Execute ...
func (command *Command) Execute(commandLine []string) error {
	for i := 0; i < len(commandLine); i++ {
		arg := commandLine[i]
		if strings.HasPrefix(arg, "--") {
			if len(arg) > len("--") {
				arg = strings.TrimPrefix(arg, "--")
				a := strings.SplitN(arg, "=", 2)
				var value *string
				if len(a) == 2 {
					value = &a[1]
				}
				if flag := command.FindFlag("", a[0]); flag != nil {
					if err := flag.Parse(value); err != nil {
						if _, ok := err.(ErrFlagNeedsValue); ok {
							if (i + 1) < len(commandLine) {
								if err = flag.Parse(&commandLine[i+1]); err != nil {
									return err
								}
								i++
							} else {
								return err
							}
						} else {
							return err
						}
					}
				} else {
					return ErrFlagNotFound("--" + arg)
				}
			} else {
				if command.Function != nil {
					return command.Function(command, commandLine[i+1:])
				}
				return nil
			}
		} else if strings.HasPrefix(arg, "-") && len(arg) > len("-") {
			flags := []rune(arg)
			for j := 1; j < len(flags); j++ {
				c := flags[j]
				if flag := command.FindFlag(string(c), ""); flag != nil {
					if err := flag.Parse(nil); err != nil {
						if _, ok := err.(ErrFlagNeedsValue); ok {
							if (i+1) < len(commandLine) && (j+1) == len(flags) {
								if err = flag.Parse(&commandLine[i+1]); err != nil {
									return err
								}
								i++
							} else {
								return err
							}
						} else {
							return err
						}
					}
				} else {
					return ErrFlagNotFound("-" + string(c))
				}
			}
		} else if subCommand := command.FindSubCommmand(arg); subCommand != nil {
			command = subCommand
		} else {
			break
		}
	}
	if command.Function != nil {
		return command.Function(command, commandLine[len(commandLine):])
	}
	return nil
}

// ErrFlagNotFound ...
type ErrFlagNotFound string

func (err ErrFlagNotFound) Error() string {
	return fmt.Sprintf("flag provided but not defined: %s", string(err))
}

// ErrFlagNeedsValue ...
type ErrFlagNeedsValue string

func (err ErrFlagNeedsValue) Error() string {
	return fmt.Sprintf("flag needs an argument: %s", string(err))
}

// CommandSet ...
type CommandSet struct {
	commands []*Command
}

// NewCommandSet ...
func NewCommandSet() *CommandSet {
	return &CommandSet{
		commands: make([]*Command, 0, 1),
	}
}

// Find ...
func (set *CommandSet) Find(name string) *Command {
	for _, command := range set.commands {
		if command.name == name {
			return command
		}
	}
	return nil
}

func (set *CommandSet) find(command *Command) (index int) {
	for index, command2 := range set.commands {
		if command == command2 {
			return index
		}
	}
	return -1
}

// Add ...
func (set *CommandSet) Add(commands ...*Command) {
	for _, command := range commands {
		if set.find(command) > -1 {
			return
		}
		set.commands = append(set.commands, command)
	}
}

// Remove ...
func (set *CommandSet) Remove(commands ...*Command) {
	for _, command := range commands {
		if i := set.find(command); i > -1 {
			set.commands = append(set.commands[:i], set.commands[i+1:]...)
		}
	}
}

// DefaultCommand ...
var DefaultCommand = NewCommand(os.Args[0])

// Execute ...
func Execute() error {
	return DefaultCommand.Execute(os.Args[1:])
}
