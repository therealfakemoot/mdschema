package mdschema

import (
	"fmt"
	"io"
	"log"

	"github.com/BurntSushi/toml"
)

type YAMLType int

const (
	YAMLString = iota
	YAMLInteger
	YAMLArray
	YAMLObject
)

func LoadSchema(r io.Reader) (Policy, error) {
	var p Policy
	raw := make(map[string]Policy)

	m, err := toml.DecodeReader(r, &raw)
	if err != nil {
		return p, fmt.Errorf("error decoding schema file: %w", err)
	}

	fmt.Println("fuck")
	log.Println("fart")
	log.Printf("%#v\n", m)
	return p, nil
}
