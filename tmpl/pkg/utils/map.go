package utils

import (
	"github.com/chtison/libgo/deepcopy"
)

type Map struct{}

func NewMap() *Map {
	return &Map{}
}

func (*Map) New() map[interface{}]interface{} {
	return make(map[interface{}]interface{})
}

func (*Map) Copy(i interface{}) interface{} {
	if i == nil {
		return make(map[interface{}]interface{})
	}
	return deepcopy.Copy(i)
}

func (*Map) Set(value interface{}, m map[interface{}]interface{}, keys ...interface{}) string {
	for _, key := range keys[:len(keys)-1] {
		if _, ok := m[key]; !ok {
			m[key] = make(map[interface{}]interface{})
		} else if _, ok := m[key].(map[interface{}]interface{}); !ok {
			return ""
		}
		m = m[key].(map[interface{}]interface{})
	}
	m[keys[len(keys)-1]] = value
	return ""

}

func (*Map) SetDefault(value interface{}, m map[interface{}]interface{}, keys ...interface{}) string {
	for _, key := range keys[:len(keys)-1] {
		if _, ok := m[key]; !ok {
			m[key] = make(map[interface{}]interface{})
		} else if _, ok := m[key].(map[interface{}]interface{}); !ok {
			return ""
		}
		m = m[key].(map[interface{}]interface{})
	}
	if _, ok := m[keys[len(keys)-1]]; !ok {
		m[keys[len(keys)-1]] = value
	}
	return ""
}
