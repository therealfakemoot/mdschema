package mdschema

type PolicyRange struct {
	Start int
	End   int
}

// Valid(interface{}) bool

type Policy struct {
	Type     YAMLType    `toml:"type"`
	Optional bool        `toml:"optional"`
	Valid    PolicyRange `toml:"valid"`
}
