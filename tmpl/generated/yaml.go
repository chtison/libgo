package generated

import (
	"io"

	"github.com/chtison/libgo/yaml"
)

type Yaml struct{}

func NewYaml() *Yaml { return &Yaml{} }

func (*Yaml) Marshal(in interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

func (*Yaml) MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	return yaml.MarshalIndent(v, prefix, indent)
}

func (*Yaml) NewDecoder(r io.Reader) *yaml.Decoder {
	return yaml.NewDecoder(r)
}

func (*Yaml) NewEncoder(w io.Writer) *yaml.Encoder {
	return yaml.NewEncoder(w)
}

func (*Yaml) ReadAll(r io.Reader, out interface{}) error {
	return yaml.ReadAll(r, out)
}

func (*Yaml) ReadFile(filename string, out interface{}) error {
	return yaml.ReadFile(filename, out)
}

func (*Yaml) Unmarshal(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}

func (*Yaml) UnmarshalStrict(in []byte, out interface{}) error {
	return yaml.UnmarshalStrict(in, out)
}
