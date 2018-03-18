package generated

import (
	"io"

	"github.com/chtison/libgo/yaml"
)

type Yaml struct{}

func NewYaml() *Yaml { return &Yaml{} }

func (*Yaml) Marshal(in interface{}) (out []byte, err error) {
	return yaml.Marshal(in)
}

func (*Yaml) NewDecoder(r io.Reader) *yaml.Decoder {
	return yaml.NewDecoder(r)
}

func (*Yaml) NewEncoder(w io.Writer) *yaml.Encoder {
	return yaml.NewEncoder(w)
}

func (*Yaml) Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}

func (*Yaml) UnmarshalStrict(in []byte, out interface{}) (err error) {
	return yaml.UnmarshalStrict(in, out)
}
