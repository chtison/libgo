package flag

import "testing"

func TestBoolValue(t *testing.T) {
	b := NewBool("v", "verbose", true)
	if b.Value() != true {
		t.Error(`b.Value() == false but true is expected.`)
	}
}

func TestBoolParseValue(t *testing.T) {
	testBoolParseValue(t, false, "1")
	testBoolParseValue(t, false, "t")
	testBoolParseValue(t, false, "T")
	testBoolParseValue(t, false, "TRUE")
	testBoolParseValue(t, false, "true")
	testBoolParseValue(t, false, "True")

	testBoolParseValue(t, true, "0")
	testBoolParseValue(t, true, "f")
	testBoolParseValue(t, true, "F")
	testBoolParseValue(t, true, "FALSE")
	testBoolParseValue(t, true, "false")
	testBoolParseValue(t, true, "False")
}

func testBoolParseValue(t *testing.T, defaultValue bool, value string) {
	b := NewBool("v", "", defaultValue)
	if err := b.Parse(&value); err != nil {
		t.Fatalf(`b.Parse(&"%s") -> error("%s")`, value, err)
	}
	if b.Value() != !defaultValue {
		t.Errorf(`b.Value() == %t but %t is expected`, b.Value(), !defaultValue)
	}
}

func TestBoolParseValueEmpty(t *testing.T) {
	b := NewBool("v", "", false)
	if err := b.Parse(nil); err != nil {
		t.Fatalf(`b.Parse(nil) -> error("%s")`, err)
	}
	if b.Value() != true {
		t.Error(`b.Value() == false but true is expected`)
	}
}

func TestBoolParseValueError(t *testing.T) {
	b := NewBool("v", "", false)
	arg := "bad"
	err := b.Parse(&arg)
	if err == nil {
		t.Error(`NewBool("v", "", false).Parse(&"bad") -> nil but error of type ErrFlagInvalidSyntax is expected.`)
	}
	if _, ok := err.(*ErrFlagInvalidSyntax); !ok {
		t.Errorf(`NewBool("v", "", false).Parse(&"bad") -> error("%s") but error is not of type ErrFlagInvalidSyntax`, err)
	}
}
