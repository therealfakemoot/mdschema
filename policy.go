package mdschema

type PolicyRange interface {
	Valid(interface{}) bool
}

type Policy struct {
	Type  YAMLType `toml: "type"`
	Valid PolicyRange
}
