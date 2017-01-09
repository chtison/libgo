package flag

import (
	"os"
	"strings"
)

// Command ...
type Command struct {
	name  string
	usage string

	localFlags  *Set
	globalFlags *Set
}

// NewCommand ...
func NewCommand(name string, function func(command *Command, args []string) error) *Command {
	command := &Command{
		localFlags:  NewSet(),
		globalFlags: NewSet(),
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

// AddLocalFlag ...
func (command *Command) AddLocalFlag(flag Flag) {
	command.localFlags.Add(flag)
}

// AddGlobalFlag ...
func (command *Command) AddGlobalFlag(flag Flag) {
	command.globalFlags.Add(flag)
}

// Execute ...
func (command *Command) Execute(commandLine []string) error {
	for _, arg := range commandLine {
		if strings.HasPrefix(arg, "--") {
			if len(arg) > len("--") {
				command.localFlags.Find("", strings.TrimPrefix(arg, "--"))
			} else {

			}
		}
	}
	return nil
}

// DefaultCommand ...
var DefaultCommand = NewCommand(os.Args[0])

// AddLocalFlag ...
func AddLocalFlag(flag Flag) {
	DefaultCommand.AddLocalFlag(flag)
}

// AddGlobalFlag ...
func AddGlobalFlag(flag Flag) {
	DefaultCommand.AddLocalFlag(flag)
}

// Execute ...
func Execute() error {
	return DefaultCommand.Execute(os.Args[1:])
}
