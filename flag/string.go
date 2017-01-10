package flag

// String ...
type String struct {
	*flag
	value string
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
	str.value = *value
	return nil
}
