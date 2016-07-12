package baseconverter

import (
	"reflect"
	"testing"
)

// *****************************************************************************
// TestBaseToDecimal

func TestBaseToDecimal(t *testing.T) {

	testBaseToDecimalValue(t, "cafe", "0123456789abcdef", 51966)
	testBaseToDecimalValue(t, "@", "@$%", 0)
	testBaseToDecimalValue(t, "10011", "01", 19)
	testBaseToDecimalValue(t, "🌛🌜🌜🌛🌛", "🌜🌛", 19)
	testBaseToDecimalValue(t, "ypnoyyyyyyyyy", "poney", 999999999)
	testBaseToDecimalValue(t, "👻🐱❌😄🍓🐶🐼👻🍓😄", "😄❌👻🐱🐶🍓😤🐼🍄", 909090909)
	testBaseToDecimalValue(t, "家家家家家好大大大家大家家好好好大家大大好大家大大", "大家好", 424242424242)
	testBaseToDecimalValue(t, "", "01", 0)
	testBaseToDecimalValue(t, "18446744073709551615", "0123456789", 18446744073709551615) // math.MaxUint64

	testBaseToDecimalError(t, "", "", ErrBaseLengthTooShort(0))
	testBaseToDecimalError(t, "", "0", ErrBaseLengthTooShort(1))
	testBaseToDecimalError(t, "", "00", ErrDuplicateCharInBase('0'))
	testBaseToDecimalError(t, "", "👻0123456789abcdef👻", ErrDuplicateCharInBase('👻'))
	testBaseToDecimalError(t, "❌", "orange", ErrCharNotInBase('❌'))
	testBaseToDecimalError(t, "18446744073709551616", "0123456789", ErrOverflow) // math.MaxUint64 + 1

}

func testBaseToDecimalValue(t *testing.T, number, inBase string, expected uint) {
	nbr, err := BaseToDecimal(number, inBase)
	if err != nil {
		t.Errorf(`BaseToDecimal("%s", "%s") returns %d and error "%s"`,
			number, inBase, nbr, err)
		return
	}
	if nbr != expected {
		t.Errorf(`BaseToDecimal("%s", "%s") returns %d but expected is %d`,
			number, inBase, nbr, expected)
		return
	}
}

func testBaseToDecimalError(t *testing.T, number, inBase string, expected error) {
	_, err := BaseToDecimal(number, inBase)
	t1, t2 := reflect.TypeOf(err), reflect.TypeOf(expected)
	if t1 != t2 || err.Error() != expected.Error() {
		t.Errorf(`BaseToDecimal("%s", "%s") returns error %s(%s) but expected is %s(%s)`,
			number, inBase, t1, err, t2, expected)
	}
}

// *****************************************************************************
// TestDecimalToBase

func TestDecimalToBase(t *testing.T) {

	testDecimalToBaseValue(t, 0, "ab", "a")
	testDecimalToBaseValue(t, 18446744073709551615, "0123456789", "18446744073709551615")
	testDecimalToBaseValue(t, 42, "🌵💲🐮", "💲💲🐮🌵")

	testDecimalToBaseError(t, 0, "", ErrBaseLengthTooShort(0))
	testDecimalToBaseError(t, 0, "x", ErrBaseLengthTooShort(1))
	testDecimalToBaseError(t, 0, "👻1234👻", ErrDuplicateCharInBase('👻'))

}

func testDecimalToBaseValue(t *testing.T, number uint, toBase, expected string) {
	nbr, err := DecimalToBase(number, toBase)
	if err != nil {
		t.Errorf(`DecimalToBase(%d, "%s") returns "%s" and error "%s"`,
			number, toBase, nbr, expected)
		return
	}
	if nbr != expected {
		t.Errorf(`DecimalToBase(%d, "%s") returns "%s" but expected is "%s"`,
			number, toBase, nbr, expected)
		return
	}
}

func testDecimalToBaseError(t *testing.T, number uint, toBase string, expected error) {
	_, err := DecimalToBase(number, toBase)
	t1, t2 := reflect.TypeOf(err), reflect.TypeOf(expected)
	if t1 != t2 || err.Error() != expected.Error() {
		t.Errorf(`BaseToDecimal(%d, "%s") returns error %s(%s) but expected is %s(%s)`,
			number, toBase, t1, err, t2, expected)
		return
	}
}
