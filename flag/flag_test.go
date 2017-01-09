package flag

import (
	"strings"
	"testing"
)

func TestFlagEmptyNames(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error(`newFlag("", "") should panic.`)
		} else if !strings.Contains(r.(string), "empty") {
			t.Errorf(`newFlag("", "") -> panic("%s")`, r.(string))
		}
	}()
	_ = newFlag("", "")
}

func TestFlagShortNameMoreThanOneRune(t *testing.T) {
	testFlagShortNameMoreThanOneRune(t, "aa")
}

func testFlagShortNameMoreThanOneRune(t *testing.T, shortName string) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(`newFlag("%s", "") should panic.`, shortName)
		} else if !strings.Contains(r.(string), "more") {
			t.Errorf(`newFlag("%s", "") -> panic("%s")`, shortName, r.(string))
		}
	}()
	_ = newFlag(shortName, "")
}

func TestFlagForbiddenCharEqual(t *testing.T) {
	testFlagForbiddenCharEqual(t, "=")
	testFlagForbiddenCharEqual(t, "==")
	testFlagForbiddenCharEqual(t, "=a")
	testFlagForbiddenCharEqual(t, "a=")
	testFlagForbiddenCharEqual(t, "x=x")
}

func testFlagForbiddenCharEqual(t *testing.T, longName string) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(`newFlag("", "%s") should panic.`, longName)
		} else if !strings.Contains(r.(string), "=") {
			t.Errorf(`newFlag("", "%s") -> panic("%s")`, longName, r.(string))
		}
	}()
	_ = newFlag("", longName)
}

func TestFlagNames(t *testing.T) {
	testFlagNames(t, "n", "name")
	testFlagNames(t, "", "name")
	testFlagNames(t, "n", "")
	testFlagNames(t, "n", "n")
	testFlagNames(t, "😙", "😙😙😙😙")
}

func testFlagNames(t *testing.T, shortName, longName string) {
	flag := newFlag(shortName, longName)
	s, l := flag.Name()
	if s != shortName || l != longName {
		t.Errorf(`newFlag("%s", "%s").Name() -> ("%s", "%s")`, shortName, longName, s, l)
	}
}

func TestFlagUsage(t *testing.T) {
	testFlagUsage(t, "")
	testFlagUsage(t, "name sets the name of the container")
}

func testFlagUsage(t *testing.T, usage string) {
	flag := newFlag("u", "usage")
	if flag.Usage() != "" {
		t.Error("Usage must be empty when not set.")
	}
	flag.SetUsage(usage)
	if flag.Usage() != usage {
		t.Errorf(`flag.Usage() != "%s"`, usage)
	}
}
