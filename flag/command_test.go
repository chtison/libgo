package flag

import (
	"strings"
	"testing"
	"time"
)

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

func TestCommandExecuteLongFlag(t *testing.T) {
	testCommandExecuteLongFlag1(t, "name", "sarah")
	testCommandExecuteLongFlag1(t, "name", "2*21=42")
	testCommandExecuteLongFlag1(t, "name", "")

	testCommandExecuteLongFlag2(t, "name", "sarah")
	testCommandExecuteLongFlag2(t, "name", "2*21=42")
	testCommandExecuteLongFlag2(t, "name", "")
}

func testCommandExecuteLongFlag1(t *testing.T, longName, value string) {
	cmd := NewCommand("test")
	cmd.LocalFlags.Add(NewString("", longName, ""))
	if err := cmd.Execute([]string{"--" + longName, value}); err != nil {
		t.Fatalf(`cmd.Execute([]string{"--%s", "%s"}]) -> error("%s")`, longName, value, err)
	}
}

func testCommandExecuteLongFlag2(t *testing.T, longName, value string) {
	cmd := NewCommand("test")
	cmd.LocalFlags.Add(NewString("", longName, ""))
	if err := cmd.Execute([]string{"--" + longName + "=" + value}); err != nil {
		t.Fatalf(`cmd.Execute([]string{"--%s", "%s"}]) -> error("%s")`, longName, value, err)
	}
}

func TestCommandExecuteErrFlagNeedsValueInsatisfied(t *testing.T) {
	cmd := NewCommand("test")

	cmd.LocalFlags.Add(NewString("", "name", ""))
	testCommandExecuteErrFlagNeedsValueInsatisfied(t, cmd, []string{"--name"})

	cmd.LocalFlags.Add(NewString("p", "", ""))
	cmd.LocalFlags.Add(NewString("q", "", ""))
	testCommandExecuteErrFlagNeedsValueInsatisfied(t, cmd, []string{"-pq"})
	testCommandExecuteErrFlagNeedsValueInsatisfied(t, cmd, []string{"-p"})
}

func testCommandExecuteErrFlagNeedsValueInsatisfied(t *testing.T, cmd *Command, args []string) {
	err := cmd.Execute(args)
	if err == nil {
		t.Fatalf(`cmd.Execute([]string{%v}) should not return nil.`, args)
	}
	if _, ok := err.(ErrFlagNeedsValue); !ok {
		t.Errorf(`error("%s") should be of type ErrFlagNeedsValue.`, err)
	}

}

func TestCommandExecuteStop(t *testing.T) {
	testCommandExecuteStop(t, []string{"--", "--name"})
	testCommandExecuteStop(t, []string{"--", ""})

	cmd := NewCommand("name")
	if err := cmd.Execute([]string{"--"}); err != nil {
		t.Errorf(`cmd.Execute([]string{"--"}) -> error("%s")`, err)
	}
}

func testCommandExecuteStop(t *testing.T, args []string) {
	cmd := NewCommand("test")
	c := make(chan []string)
	cmd.Function = func(cmd *Command, args []string) error {
		c <- args
		return nil
	}
	go cmd.Execute(args)
	select {
	case args2 := <-c:
		for i, arg := range args[1:] {
			if arg != args2[i] {
				t.Errorf(`%v != %v`, args2, args[1:])
			}
		}
	case <-time.After(1 * time.Second):
		t.Error(`Function Not Called !`)
	}
}

func TestCommandExecuteErrFlagNotFound(t *testing.T) {
	testCommandExecuteErrFlagNotFound(t, []string{"--name"})
	testCommandExecuteErrFlagNotFound(t, []string{"--="})
	testCommandExecuteErrFlagNotFound(t, []string{"-x"})
}

func testCommandExecuteErrFlagNotFound(t *testing.T, args []string) {
	cmd := NewCommand("name")
	err := cmd.Execute(args)
	if err == nil {
		t.Fatalf(`cmd.Execute(%v) should return error of type ErrFlagNotFound`, args)
	}
	if _, ok := err.(ErrFlagNotFound); !ok {
		t.Errorf(`error("%s") is not of type ErrFlagNotFound`, err.Error())
	}
}

func TestCommandExecuteShortFlag(t *testing.T) {
	testCommandExecuteShortFlag(t, "n", "alexandre")
	testCommandExecuteShortFlag(t, "n", "")
}

func testCommandExecuteShortFlag(t *testing.T, shortName, value string) {
	cmd := NewCommand("name")
	cmd.LocalFlags.Add(NewString(shortName, "", ""))
	if err := cmd.Execute([]string{"-" + shortName, value}); err != nil {
		t.Fatalf(`cmd.Execute([]string{"-%s", "%s"}) -> error("%s")`, shortName, value, err)
	}
	v := cmd.FindFlag(shortName, "").(*String).Value()
	if v != value {
		t.Errorf(`cmd.FindFlag("%s", "").(*String).Value() == "%s" but "%s" is expected`, shortName, v, value)
	}

}

func TestCommandExecuteEmptyParams(t *testing.T) {
	cmd := NewCommand("test")
	if err := cmd.Execute([]string{}); err != nil {
		t.Errorf(`cmd.Execute([]string{}) -> error("%s")`, err)
	}

	c := make(chan bool)
	cmd.Function = func(cmd *Command, args []string) error {
		c <- true
		return nil
	}
	go cmd.Execute([]string{})
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		t.Error(`Function not call !`)
	}
}

func TestCommandExecuteSubCommand(t *testing.T) {
	cmd := NewCommand("docker")
	cmd1 := NewCommand("run")
	cmd2 := NewCommand("build")
	cmd3 := NewCommand("ps")
	c := make(chan bool)
	cmd2.Function = func(cmd *Command, args []string) error {
		c <- true
		return nil
	}
	cmd.AddSubCommand(cmd1, cmd2, cmd3)
	go cmd.Execute([]string{"build", "."})
	select {
	case <-c:
	case <-time.After(time.Second):
		t.Error(`Function not called !`)
	}
}

func TestErrFlagNotFound(t *testing.T) {
	if !strings.HasSuffix(ErrFlagNotFound("--name").Error(), "--name") {
		t.Error(`strings.HasSuffix(ErrFlagNotFound("--name").Error(), "--name") == false`)
	}
}
