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

// ParseValue ...
func (str *String) ParseValue(value string) error {
	str.value = value
	return nil
}
