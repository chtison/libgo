package flag

import "testing"

func TestStringValue(t *testing.T) {
	testStringValue(t, "")
	testStringValue(t, "a")
	testStringValue(t, "reverent_euclid")
	testStringValue(t, "=")
}

func testStringValue(t *testing.T, value string) {
	flag := NewString("v", "value", "value")
	if flag.Value() != "value" {
		t.Errorf(`NewValue("value", "value", "value").Value() == "%s"`, flag.Value())
	}
	if err := flag.ParseValue(value); err != nil {
		t.Errorf(`NewValue("value", "value", "value").ParseValue("%s") == "%s"`, value, err)
		return
	}
	if flag.Value() != value {
		t.Errorf(`NewValue("value", "value", "value").ParseValue("%s").Value() == "%s"`, value, flag.Value())
	}
}
