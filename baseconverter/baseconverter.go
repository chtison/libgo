package baseconverter

import (
	"container/list"
	"fmt"
)

// DefaultBaseConverter is the default BaseConverter used by some functions of this package.
var DefaultBaseConverter = &BaseConverter{
	StopOnOverflow: false,
}

// ErrOverflow error is an overflow flag.
var ErrOverflow = fmt.Errorf("uint overflow")

type (
	// ErrBaseLengthTooShort is thrown when the length of a base is less than 2.
	//
	// The underlying uint is the length of the base.
	ErrBaseLengthTooShort uint
	// ErrDuplicateCharInBase is thrown when there is (at least) one duplicate character in base.
	//
	// The underlying rune is the found duplicated character.
	ErrDuplicateCharInBase rune
	// ErrCharNotInBase is thrown when there is no correspondence for a character in a base.
	//
	// The underlying rune is the homeless character.
	ErrCharNotInBase rune
)

func (err ErrBaseLengthTooShort) Error() string {
	return fmt.Sprintf("base length too short: %d", uint(err))
}

func (err ErrDuplicateCharInBase) Error() string {
	return fmt.Sprintf("duplicate character in base: '%c'", rune(err))
}

func (err ErrCharNotInBase) Error() string {
	return fmt.Sprintf("number's character not in base: '%c'", rune(err))
}

// BaseConverter is used to perform base conversions.
type BaseConverter struct {
	// StopOnOverflow determines if the execution should stop if an
	// ErrOverflow occurs (but does not prevent the error from being returned).
	StopOnOverflow bool
}

/*
BaseToDecimal converts a number in any base to a number in base decimal.

All runes in the number string must be present in the inBase string.

An empty number string returns 0 as result and no error.

BaseToDecimal can return the following errors:
	ErrBaseLengthTooShort(baseLength)     // if len(inBase) < 2
	ErrDuplicateCharInBase(duplicateRune) // if inBase[i] == inBase[j] with i != j
	ErrCharNotInBase(missingRune)         // if number[i] is not in inBase
	ErrOverflow                           // if an overflow has occured on result

BaseToDecimal will stop its execution when ErrOverflow occurs if its underlying
BaseConverter's StopOnOverflow property is set to true, otherwise it continues
normally its execution, but still return the ErrOverflow error however.
*/
func (baseConverter *BaseConverter) BaseToDecimal(number string, inBase string) (result uint, err error) {
	base := []rune(inBase)
	if err = checkBase(base); err != nil {
		return
	}
	for _, c := range number {
		i := indexRune(base, c)
		if i < 0 {
			return 0, ErrCharNotInBase(c)
		}
		tmp := result*uint(len(base)) + uint(i)
		if tmp < result {
			err = ErrOverflow
			if baseConverter.StopOnOverflow {
				return 0, ErrOverflow
			}
		}
		result = tmp
	}
	return
}

// BaseToDecimal binds to DefaultBaseConverter.BaseToDecimal(:,:) : (:,:)
func BaseToDecimal(number string, inBase string) (result uint, err error) {
	return DefaultBaseConverter.BaseToDecimal(number, inBase)
}

/*
DecimalToBase converts a number in decimal base to the the specified base.

DecimalToBase can return the following errors:
	ErrBaseLengthTooShort(baseLength)     // if len(toBase) < 2
	ErrDuplicateCharInBase(duplicateRune) // if toBase[i] == toBase[j] with i != j
*/
func (baseConverter *BaseConverter) DecimalToBase(number uint, toBase string) (result string, err error) {
	base := []rune(toBase)
	if err = checkBase(base); err != nil {
		return
	}
	if number == 0 {
		return string(base[0]), nil
	}
	baseLen := uint(len(base))
	l := list.New()
	for ; number > 0; number /= baseLen {
		l.PushFront(base[number%baseLen])
	}
	a := make([]rune, l.Len())
	for e, i := l.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		a[i] = e.Value.(rune)
	}
	return string(a), nil
}

// DecimalToBase binds to DefaultBaseConverter.DecimalToBase(:,:) : (:, :)
func DecimalToBase(number uint, toBase string) (result string, err error) {
	return DefaultBaseConverter.DecimalToBase(number, toBase)
}

/*
BaseToBase has the following definition:
	func BaseToBase(number string, inBase string, toBase string) (result string, err error) {
		nbr, err := BaseToDecimal(number, inBase)
		if err != nil || err == ErrOverflow {
			return
		}
		result, err = DecimalToBase(nbr, toBase)
		return
	}
*/
func (baseConverter *BaseConverter) BaseToBase(number string, inBase string, toBase string) (result string, err error) {
	nbr, err := baseConverter.BaseToDecimal(number, inBase)
	if err != nil || err == ErrOverflow {
		return
	}
	result, err = DecimalToBase(nbr, toBase)
	return
}

// BaseToBase binds to DefaultBaseConverter.BaseToBase(:,:) : (:,:)
func BaseToBase(number string, inBase string, toBase string) (result string, err error) {
	return DefaultBaseConverter.BaseToBase(number, inBase, toBase)
}

// Checks if a base has a length less than 2 or any duplicate character.
func checkBase(base []rune) error {
	if len(base) < 2 {
		return ErrBaseLengthTooShort(len(base))
	}
	for i := range base {
		for j := range base[i+1:] {
			if base[i] == base[i+1+j] {
				return ErrDuplicateCharInBase(base[i])
			}
		}
	}
	return nil
}

// Same behavior than strings.IndexRune(:,:) but takes a slice of rune instead of a string.
func indexRune(s []rune, c rune) int {
	for i := range s {
		if s[i] == c {
			return i
		}
	}
	return -1
}
