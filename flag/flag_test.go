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

func TestFlagErrFlagNeedsValue(t *testing.T) {
	testFlagErrFlagNeedsValue(t, "", "long", "--long")
	testFlagErrFlagNeedsValue(t, "l", "", "-l")
	testFlagErrFlagNeedsValue(t, "l", "long", "--long")
}

func testFlagErrFlagNeedsValue(t *testing.T, shortName, longName, expected string) {
	flag := newFlag(shortName, longName)
	if strings.HasSuffix(flag.errFlagNeedsValue().Error(), expected) == false {
		t.Errorf(`newFlag("l", "long").errFlagNeedsValue().Error() == "%s" but expected is "%s"`, flag.errFlagNeedsValue().Error(), expected)
	}
}

func TestFlagSetFind(t *testing.T) {
	set := NewFlagSet()
	set.Add(NewString("n", "name", ""))
	set.Add(NewString("i", "interface", "bridge"))
	set.Add(NewString("u", "user", "root"))

	testFlagSetFind(t, set, "n", "", 0)
	testFlagSetFind(t, set, "n", "name", 0)
	testFlagSetFind(t, set, "", "name", 0)
	testFlagSetFind(t, set, "u", "user", 2)
	testFlagSetFind(t, set, "", "user", 2)
	testFlagSetFind(t, set, "u", "", 2)

	s, l := set.Find("n", "name").Name()
	if s != "n" || l != "name" {
		t.Errorf(`set.Find("%s", "%s").Name() -> ("%s", "%s") but expeted is ("%s", "%s")`,
			"n", "name", s, l, "n", "name")
	}
}

func testFlagSetFind(t *testing.T, set *FlagSet, shortName, longName string, expectedIndex int) {
	index, flag := set.find(shortName, longName)
	if index != expectedIndex {
		t.Errorf(`set.find("%s", "%s") returns index %d but expected is %d`, shortName, longName, index, expectedIndex)
	}
	s, l := flag.Name()
	if s != shortName && l != longName {
		t.Errorf(`_, f := set.find("%s", "%s"); f.Name() -> ("%s", "%s") but expeted is ("%s", "%s")`,
			shortName, longName, s, l, shortName, longName)
	}
}

func TestFlagSetRemove(t *testing.T) {
	set := NewFlagSet()
	set.Add(NewString("n", "name", ""))
	set.Add(NewString("i", "interface", "bridge"))
	set.Add(NewString("u", "user", "root"))

	testFlagSetRemove(t, set, "i", "interface")
	testFlagSetRemove(t, set, "n", "name")
	set.Remove("u", "")
}

func testFlagSetRemove(t *testing.T, set *FlagSet, shortName, longName string) {
	set.Remove("", longName)
	if set.Find(shortName, longName) != nil {
		s, l := set.Find(shortName, longName).Name()
		t.Errorf(`set.Find("%s", "%s") -> flag.Name() -> ("%s", "%s") but should have returned nil`, shortName, longName, s, l)
	}
}
