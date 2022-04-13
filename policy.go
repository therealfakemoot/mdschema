package mdschema

type IntValidator struct {
	Start, End, Step int
}

// Valid(interface{}) bool

type Policy struct {
	Type     YAMLType    `toml:"type"`
	Optional bool        `toml:"optional"`
	Valid    PolicyRange `toml:"valid"`
}
