package flag

// String ...
type String struct {
	*flag
	value     string
	Validator func(str *String, value string) (newValue string, err error)
}

// NewString ...
func NewString(shortName, longName, defaultValue string) *String {
	return &String{
		flag:  newFlag(shortName, longName),
		value: defaultValue,
	}
}

// Value ...
func (str *String) Value() string {
	return str.value
}

// Parse ...
func (str *String) Parse(value *string) error {
	if value == nil {
		return str.errFlagNeedsValue()
	}
	if str.Validator != nil {
		newValue, err := str.Validator(str, *value)
		if err != nil {
			return newErrFlagInvalidSyntax(*value, err)
		}
		str.value = newValue
		return nil
	}
	str.value = *value
	return nil
}
