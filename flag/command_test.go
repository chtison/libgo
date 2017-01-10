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

func TestCommandFindFlag(t *testing.T) {

}
