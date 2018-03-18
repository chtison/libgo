package generated

import (
	"bytes"
	"encoding/json"
	"io"
)

type Json struct{}

func NewJson() *Json { return &Json{} }

func (*Json) Compact(dst *bytes.Buffer, src []byte) error {
	return json.Compact(dst, src)
}

func (*Json) HTMLEscape(dst *bytes.Buffer, src []byte) {
	json.HTMLEscape(dst, src)
}

func (*Json) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}

func (*Json) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*Json) MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (*Json) NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

func (*Json) NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func (*Json) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (*Json) Valid(data []byte) bool {
	return json.Valid(data)
}
