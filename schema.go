package mdschema

import (
	"io"
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

	return p, nil
}
