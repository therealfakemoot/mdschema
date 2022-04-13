package mdschema

import (
	"fmt"
	"io"
	// "log"

	"github.com/BurntSushi/toml"
)

type YAMLType int

const (
	YAMLString = iota
	YAMLInteger
	YAMLArray
	YAMLObject
)

func LoadSchema(r io.Reader) (map[string]Policy, error) {
	raw := make(map[string]Policy)

	_, err := toml.DecodeReader(r, &raw)
	if err != nil {
		return raw, fmt.Errorf("error decoding schema file: %w", err)
	}

	return raw, nil
}
