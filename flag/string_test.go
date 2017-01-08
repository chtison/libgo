package flag

import (
	"strings"
	"testing"
)

func TestStringEmptyNames(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error(`NewString("", "", "") should panic.`)
		} else if !strings.Contains(r.(string), "empty") {
			t.Errorf(`NewString("", "", "") -> panic("%s")`, r.(string))
		}
	}()
	_ = NewString("", "", "")
}

func TestStringShortNameMoreThanOneRune(t *testing.T) {
	testStringShortNameMoreThanOneRune(t, "aa")
}

func testStringShortNameMoreThanOneRune(t *testing.T, shortName string) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(`NewString("%s", "", "") should panic.`, shortName)
		} else if !strings.Contains(r.(string), "more") {
			t.Errorf(`NewString("%s", "", "") -> panic("%s")`, shortName, r.(string))
		}
	}()
	_ = NewString(shortName, "", "")
}

func TestStringForbiddenCharEqual(t *testing.T) {
	testStringForbiddenCharEqual(t, "=")
	testStringForbiddenCharEqual(t, "==")
	testStringForbiddenCharEqual(t, "=a")
	testStringForbiddenCharEqual(t, "a=")
	testStringForbiddenCharEqual(t, "x=x")
}

func testStringForbiddenCharEqual(t *testing.T, longName string) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf(`NewString("", "%s", "") should panic.`, longName)
		} else if !strings.Contains(r.(string), "=") {
			t.Errorf(`NewString("", "%s", "") -> panic("%s")`, longName, r.(string))
		}
	}()
	_ = NewString("", longName, "")
}

func TestStringNames(t *testing.T) {
	testStringNames(t, "n", "name")
	testStringNames(t, "", "name")
	testStringNames(t, "n", "")
	testStringNames(t, "n", "n")
	testStringNames(t, "😙", "😙😙😙😙")
}

func testStringNames(t *testing.T, shortName, longName string) {
	flag := NewString(shortName, longName, "")
	s, l := flag.Name()
	if s != shortName || l != longName {
		t.Errorf(`NewString("%s", "%s").Name() -> ("%s", "%s")`, shortName, longName, s, l)
	}
}

func TestStringUsage(t *testing.T) {
	testStringUsage(t, "")
	testStringUsage(t, "name sets the name of the container")
}

func testStringUsage(t *testing.T, usage string) {
	flag := NewString("u", "usage", "usage")
	if flag.Usage() != "" {
		t.Error("Usage must be empty when not set.")
	}
	flag.SetUsage(usage)
	if flag.Usage() != usage {
		t.Errorf(`flag.Usage() != "%s"`, usage)
	}
}

func TestStringValue(t *testing.T) {
	testStringValue(t, "")
	testStringValue(t, "a")
	testStringValue(t, "reverent_euclid")
	testStringValue(t, "=")
}

func testStringValue(t *testing.T, value string) {
	flag := NewString("v", "value", "value")
	if flag.Value() != "value" {
		t.Errorf(`NewValue("value", "value", "value").Value() == "%s"`, flag.Value())
	}
	if err := flag.ParseValue(value); err != nil {
		t.Errorf(`NewValue("value", "value", "value").ParseValue("%s") == "%s"`, value, err)
		return
	}
	if flag.Value() != value {
		t.Errorf(`NewValue("value", "value", "value").ParseValue("%s").Value() == "%s"`, value, flag.Value())
	}
}
