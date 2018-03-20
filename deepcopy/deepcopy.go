package deepcopy

import (
	"fmt"
	"reflect"
)

func Copy(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	copy := CopyValue(reflect.ValueOf(i))
	if !copy.IsValid() {
		return nil
	}
	return copy.Interface()
}

func CopyValue(value reflect.Value) reflect.Value {
	if !value.IsValid() {
		return value
	}
	copy := reflect.New(value.Type()).Elem()
	switch value.Kind() {
	case reflect.Invalid:
		panic("should not happend")
	case reflect.Array, reflect.Chan, reflect.Func, reflect.Ptr, reflect.Struct, reflect.UnsafePointer:
		panic(fmt.Sprintf("kind %s not supported", value.Kind()))
	case reflect.Interface:
		if value.IsNil() {
			break
		}
		copy.Set(CopyValue(reflect.ValueOf(value.Interface())))
	case reflect.Map:
		if value.IsNil() {
			break
		}
		copy.Set(reflect.MakeMapWithSize(value.Type(), value.Len()))
		for _, key := range value.MapKeys() {
			copy.SetMapIndex(CopyValue(key), CopyValue(value.MapIndex(key)))
		}
	case reflect.Slice:
		if value.IsNil() {
			break
		}
		copy.Set(reflect.MakeSlice(value.Type(), value.Len(), value.Cap()))
		for i := 0; i < value.Len(); i++ {
			copy.Index(i).Set(CopyValue(value.Index(i)))
		}
	default:
		copy.Set(value)
	}
	return copy
}
