package yaml

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"

	"github.com/chtison/libgo/fmt"
	"github.com/go-yaml/yaml"
)

// TODO: maybe fix list indentation
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	out, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	builder := fmt.NewBuilder()
	first := true
	for scanner.Scan() {
		if first {
			first = false
		} else {
			builder.Print(prefix)
		}
		text := scanner.Text()
		for i, c := range text {
			if c != ' ' {
				text = strings.Replace(text, "  ", indent, i / 2)
				break
			}
		}
		builder.Println(text)
	}
	return []byte(builder.String()), scanner.Err()
}

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
