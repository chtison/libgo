package flag

import (
	"errors"
	"testing"
)

func TestStringValue(t *testing.T) {
	testStringParse(t, "")
	testStringParse(t, "a")
	testStringParse(t, "reverent_euclid")
	testStringParse(t, "=")
	testStringParse(t, "😄😁")
}

func testStringParse(t *testing.T, value string) {
	flag := NewString("v", "value", "value")
	if flag.Value() != "value" {
		t.Errorf(`NewString("v", "value", "value").Value() == "%s"`, flag.Value())
	}
	if err := flag.Parse(&value); err != nil {
		t.Errorf(`NewString("v", "value", "value").ParseValue(&"%s") -> error("%s")`, value, err)
		return
	}
	if flag.Value() != value {
		t.Errorf(`f := NewString("v", "value", "value"); f.ParseValue("%s"); f.Value() == "%s"`, value, flag.Value())
	}
}

func TestStringParseNil(t *testing.T) {
	err := NewString("", "error", "asd").Parse(nil)
	if err == nil {
		t.Error(`NewString("", "error", "asd").Parse(nil) -> nil but should return error of type ErrFlagNeedsValue`)
		return
	}
	if _, ok := err.(ErrFlagNeedsValue); !ok {
		t.Errorf(`NewString("", "error", "asd").Parse(nil) -> error("%s") but should return error of type ErrFlagNeedsValue`, err)
	}
}

func TestStringValidator(t *testing.T) {
	s := NewString("", "1", "")
	s.Validator = func(str *String, value string) (newValue string, err error) {
		switch value {
		case "1":
			return "one", nil
		case "2":
			return "two", nil
		}
		return "", errors.New("bad value")
	}
	arg := "1"
	if err := s.Parse(&arg); err != nil {
		t.Fatalf(`s.Parse(&"%s") -> error("%s")`, arg, err)
	}
	if s.Value() != "one" {
		t.Fatalf(`s.Value() == "%s" but "one" is expected`, s.Value())
	}

	arg = "wrong"
	if err := s.Parse(&arg); err == nil {
		t.Fatalf(`s.Parse(&"%s") -> nil but error of type ErrFlagInvalidSyntax is expected`, arg)
	}
}
