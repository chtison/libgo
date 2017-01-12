package flag

import "strconv"

// Bool ...
type Bool struct {
	*flag
	value bool
}

// NewBool ...
func NewBool(shortName, longName string, defaultValue bool) *Bool {
	return &Bool{
		flag:  newFlag(shortName, longName),
		value: defaultValue,
	}
}

// Value ...
func (b *Bool) Value() bool {
	return b.value
}

// Parse ...
func (b *Bool) Parse(value *string) error {
	if value == nil {
		b.value = true
		return nil
	}
	newValue, err := strconv.ParseBool(*value)
	if err != nil {
		return newErrFlagInvalidSyntax(*value, err)
	}
	b.value = newValue
	return nil
}
