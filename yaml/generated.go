package yaml

import "github.com/go-yaml/yaml"

var (
	Marshal         = yaml.Marshal
	NewDecoder      = yaml.NewDecoder
	NewEncoder      = yaml.NewEncoder
	Unmarshal       = yaml.Unmarshal
	UnmarshalStrict = yaml.UnmarshalStrict
)

type (
	Decoder     = yaml.Decoder
	Encoder     = yaml.Encoder
	IsZeroer    = yaml.IsZeroer
	MapItem     = yaml.MapItem
	MapSlice    = yaml.MapSlice
	Marshaler   = yaml.Marshaler
	TypeError   = yaml.TypeError
	Unmarshaler = yaml.Unmarshaler
)
