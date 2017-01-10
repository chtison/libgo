package flag

import "testing"

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
