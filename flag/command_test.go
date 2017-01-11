package flag

import "testing"

func TestCommandName(t *testing.T) {
	cmd := NewCommand("test")
	if cmd.Name() != "test" {
		t.Errorf(`NewCommand("test").Name() -> "%s"`, cmd.Name())
	}
}

func TestCommandSetEmptyName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(`NewCommand("test").SetName("") should panic.`)
		}
	}()
	NewCommand("test").SetName("")
}

func TestCommandSetName(t *testing.T) {
	name := "test"
	cmd := NewCommand("name")
	cmd.SetName(name)
	if cmd.Name() != name {
		t.Errorf(`NewCommand("name").SetName("%s").Name() != "%s"`, name, cmd.Name())
	}
}

func TestCommandUsage(t *testing.T) {
	cmd := NewCommand("name")
	if cmd.Usage() != "" {
		t.Errorf(`NewCommand("name").Usage() -> "%s" but expected is empty string`, cmd.Usage())
	}
	cmd.SetUsage("usage")
	if cmd.Usage() != "usage" {
		t.Errorf(`NewCommand("name").SetUsage("usage").Usage() -> "%s"`, cmd.Usage())
	}
}

func TestCommandParent(t *testing.T) {
	cmd1 := NewCommand("parent")
	cmd2 := NewCommand("child")
	cmd1.AddSubCommand(cmd2)
	if cmd1.Parent() != nil {
		t.Errorf(`cmd1.Parent().Name() -> %s but parent should be nil`, cmd1.Parent().Name())
	}
	if cmd2.Parent() == nil {
		t.Error(`cmd2.Parent() == nil but cmd1 is exptected`)
	} else if cmd2.Parent().Name() != "parent" {
		t.Errorf(`cmd2.Parent().Name() == "%s" but "parent" is expected`, cmd2.Parent().Name())
	}
}

func TestCommandFindFlag(t *testing.T) {
	cmd1 := NewCommand("name")
	cmd1.GlobalFlags.Add(NewString("n", "name", "none"))
	cmd2 := NewCommand("name")
	cmd1.AddSubCommand(cmd2)
	testCommandFindFlag(t, cmd2, "n", "name", "none")
	cmd2.LocalFlags.Add(NewString("n", "name", "franc"))
	testCommandFindFlag(t, cmd2, "n", "name", "franc")

	if flag := cmd2.FindFlag("t", "type"); flag != nil {
		s, l := flag.Name()
		t.Errorf(`cmd2.FindFlag("t", "type").Name() -> ("%s", "%s")`, s, l)
	}
}

func testCommandFindFlag(t *testing.T, command *Command, shortName, longName, value string) {
	flag := command.FindFlag(shortName, longName)
	if flag == nil {
		t.Fatalf(`command.FindFlag("%s", "%s") -> nil but a valid flag is expected.`, shortName, longName)
	}
	if s, l := flag.Name(); s != shortName || l != longName {
		t.Errorf(`command.FindFlag("%s", "%s").Name() -> ("%s", "%s")`, shortName, longName, s, l)
	}
	s, ok := flag.(*String)
	if !ok {
		t.Fatal(`flag.(*String) returns false`)
	}
	if s.Value() != value {
		t.Errorf(`s.Value() == "%s" but "%s" is expected`, s.Value(), value)
	}
}

func TestCommandFindCommand(t *testing.T) {
	cmd1 := NewCommand("grant-parent")
	cmd2 := NewCommand("parent")
	cmd3 := NewCommand("child")
	cmd1.AddSubCommand(cmd2)
	cmd2.AddSubCommand(cmd3)

	if cmd := cmd1.FindSubCommmand("parent"); cmd == nil {
		t.Error(`cmd1.FindSubCommand("parent") -> nil but a valid *Command is expected`)
	}
	if cmd := cmd1.FindSubCommmand("child"); cmd != nil {
		t.Errorf(`cmd1.FindSubCommand("child").Name() == "%s" but nil is expected`, cmd1.FindSubCommmand("child").Name())
	}
}

func TestCommandAddSubCommand(t *testing.T) {
	cmd1 := NewCommand("cmd1")
	cmd2 := NewCommand("cmd2")
	cmd3 := NewCommand("cmd3")
	cmd1.AddSubCommand(cmd2)
	cmd3.AddSubCommand(cmd2)
	if cmd1.FindSubCommmand("cmd2") != nil {
		t.Error(`cmd1.FindSubCommand("cmd2") should return nil`)
	}
	if cmd3.FindSubCommmand("cmd2") == nil {
		t.Error(`cmd3.FindSubCommand("cmd2") should not return nil`)
	}
}

func TestCommandRemoveSubCommand(t *testing.T) {
	cmd1 := NewCommand("cmd1")
	cmd2 := NewCommand("cmd2")
	cmd3 := NewCommand("cmd3")
	cmd1.AddSubCommand(cmd2)
	cmd2.AddSubCommand(cmd3)
	cmd1.RemoveSubCommand(cmd3)
	if cmd3.Parent() == nil {
		t.Error(`cmd3.Parent() should not be equal to nil`)
	}
}

func TestCommandSetAdd(t *testing.T) {
	set := NewCommandSet()
	cmd1 := NewCommand("name")
	set.Add(cmd1)
	set.Add(cmd1)
	if len(set.commands) != 1 {
		t.Errorf(`len(set.commands) == %d but 1 is expected`, len(set.commands))
	}
}

func TestExecute(t *testing.T) {
	_ = Execute()
}
