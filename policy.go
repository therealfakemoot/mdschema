package mdschema

type PolicyRange interface {
	Valid(interface{}) bool
}

type Policy struct {
	Key   string
	Type  YAMLType
	Range PolicyRange
}
