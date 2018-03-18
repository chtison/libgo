package yaml

import (
	"io"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func ReadAll(r io.Reader, out interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, out)
}

func ReadFile(filename string, out interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, out)
}
