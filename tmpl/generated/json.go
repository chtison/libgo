package generated

import (
	"bytes"
	"encoding/json"
	"io"
)

type Json struct{}

func NewJson() *Json { return &Json{} }

func (_ *Json) Compact(dst *bytes.Buffer, src []byte) error {
	return json.Compact(dst, src)
}

func (_ *Json) HTMLEscape(dst *bytes.Buffer, src []byte) {
	json.HTMLEscape(dst, src)
}

func (_ *Json) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}

func (_ *Json) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (_ *Json) MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (_ *Json) NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

func (_ *Json) NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func (_ *Json) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (_ *Json) Valid(data []byte) bool {
	return json.Valid(data)
}
