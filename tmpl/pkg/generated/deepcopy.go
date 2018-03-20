package generated

import (
	"reflect"

	"github.com/chtison/libgo/deepcopy"
)

type Deepcopy struct{}

func NewDeepcopy() *Deepcopy { return &Deepcopy{} }

func (*Deepcopy) Copy(i interface{}) interface{} {
	return deepcopy.Copy(i)
}

func (*Deepcopy) CopyValue(value reflect.Value) reflect.Value {
	return deepcopy.CopyValue(value)
}
