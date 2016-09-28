package pwgen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	var str string

	// Test n == 0
	str = Generate(0, "charset")
	if len(str) != 0 {
		t.Errorf(`Generate(0, "charset") returns %q`, str)
	}

	// Test (n > 0) == len(str)
	// Warning: these tests assert length of string type, which is ok with single byte runes.
	for i := uint(0); i < 10; i++ {
		str = Generate(i, CharsetNumeric+CharsetLowercase)
		if len(str) != int(i) {
			t.Errorf(`Generate(%d, CharsetNumeric + CharsetLowercase) returns %q`, i, str)
		}
	}

	// Test len(str) == 0 and no panic
	str = Generate(0, "")
	if len(str) != 0 {
		t.Errorf(`Generate(0, "") returns %q`, str)
	}

	testPanic(t)
}

func testPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(`Generate(42, "") didn't panic as expected`)
		}
	}()
	_ = Generate(42, "")
}
