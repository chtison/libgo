package utils

import "reflect"

func IsType(i interface{}, typeString string) bool {
	if t := reflect.TypeOf(i); t != nil {
		return t.String() == typeString
	}
	return false
}
